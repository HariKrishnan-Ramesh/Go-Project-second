package managers

import (
	"errors"
	"fmt"
	"io"
	"main/database"
	"main/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AdminManager interface {
	UploadLogo(src io.ReadCloser, fileHeader *multipart.FileHeader) (map[string]string, error)
	GetLogo() (map[string]string, error)
	IsAdmin(userID uint) (bool, error)
	UploadImageToS3(src io.ReadCloser, fileHeader *multipart.FileHeader) (string, error)
}

type adminManager struct {
	//dbclient
}

func NewAdminManager() AdminManager {
	return &adminManager{}
}

func (adminManager *adminManager) UploadLogo(src io.ReadCloser, fileHeader *multipart.FileHeader) (map[string]string, error) {

	logoURL, err := adminManager.UploadImageToS3(src, fileHeader)

	if err != nil {
		return nil, fmt.Errorf("failed to upload Logo to S3: %v", err)
	}

	var config models.Configuration
	result := database.DB.First(&config)

	if result.Error != nil {
		config = models.Configuration{
			LogoURL: logoURL,
		}
		result = database.DB.Create(&config)
		if result.Error != nil {
			return nil, fmt.Errorf("failed to save configuration to DB: %v", result.Error)
		}
	} else {
		config.LogoURL = logoURL
		result = database.DB.Save(&config)

		if result.Error != nil {
			return nil, fmt.Errorf("failed to save configuration to DB: %v", result.Error)
		}
	}
	return map[string]string{"logoURL": logoURL}, nil
}

func (adminManager *adminManager) GetLogo() (map[string]string, error) {
	var config models.Configuration
	result := database.DB.First(&config)

	if result.Error != nil {
		return nil, fmt.Errorf("no logo found. Please upload a logo: %v", result.Error)
	}
	return map[string]string{"logoURL": config.LogoURL}, nil
}

func (adminManager *adminManager) IsAdmin(userID uint) (bool, error) {
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf("user not found: %w", result.Error)
		}
		return false, fmt.Errorf("error finding user: %w", result.Error)
	}
	if user.IsAdmin == nil {
		return false, nil
	}
	return *user.IsAdmin, nil
}

func (adminManager *adminManager) UploadImageToS3(src io.ReadCloser, fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	fileName := strings.ReplaceAll(fileHeader.Filename, ext, "")
	uniqueFileName := fmt.Sprintf("%s-%d%s", fileName, time.Now().UnixNano(), ext)

	// Ensure uploads directory exists
	uploadPath := "./uploads/"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create upload directory: %w", err)
		}
	}

	dst, err := os.Create(uploadPath + uniqueFileName)
	if err != nil {
		return "", fmt.Errorf("failed to create the file for writing: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save the file: %w", err)
	}

	imageURL := "http://localhost:8080/uploads/" + uniqueFileName
	return imageURL, nil
}
