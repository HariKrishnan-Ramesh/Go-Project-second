package managers

import (
	"errors"
	"fmt"
	"io"
	"main/common"
	"main/database"
	"main/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type HeroBannerManager interface {
	CreateHeroBanner(input common.HeroBannerCreationInput) (common.HeroBannerResponse, error)
	GetHeroBanner(heroBannerId uint) (common.HeroBannerResponse, error)
	AddImages(herobannerid uint, imgurl []string) ([]common.HeroBannerImageResponse, error)
}

type heroBannerManager struct {
	//dbclient
}

func NewHeroBannerManager() HeroBannerManager {
	return &heroBannerManager{}
}

func (manager *heroBannerManager) CreateHeroBanner(input common.HeroBannerCreationInput) (common.HeroBannerResponse, error) {
	heroBanner := models.HeroBanner{
		Title:       input.Title,
		Description: input.Description,
	}

	result := database.DB.Create(&heroBanner)
	if result.Error != nil {
		return common.HeroBannerResponse{}, fmt.Errorf("failed to create hero banner: %w", result.Error)
	}
	banner := common.HeroBannerResponse{
		Id:          heroBanner.Id,
		Title:       heroBanner.Title,
		Description: heroBanner.Description,
	}

	return banner, nil
}

func (manager *heroBannerManager) GetHeroBanner(heroBannerId uint) (common.HeroBannerResponse, error) {
	var heroBanner models.HeroBanner
	result := database.DB.Preload("Images").First(&heroBanner, heroBannerId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return common.HeroBannerResponse{}, fmt.Errorf("hero banner not found: %w", result.Error)
		}
		return common.HeroBannerResponse{}, fmt.Errorf("failed to retrieve hero banner: %w", result.Error)
	}

	var herobannerresponse = common.HeroBannerResponse{
		Id:          heroBanner.Id,
		Title:       heroBanner.Title,
		Description: heroBanner.Description,
		Image:       []common.HeroBannerImageResponse{},
	}

	for _, image := range heroBanner.Images {
		var heroimage = common.HeroBannerImageResponse{
			ImageURL: image.ImageURL,
		}
		herobannerresponse.Image = append(herobannerresponse.Image, heroimage)
	}
	return herobannerresponse, nil
}

func (managers *heroBannerManager) AddImages(herobannerid uint, imgurl []string) ([]common.HeroBannerImageResponse, error) {
	var image []models.HeroBannerImage
	for _, imageURL := range imgurl {
		var HeroBannerImage = models.HeroBannerImage{
			HeroBannerID: herobannerid,
			ImageURL:     imageURL,
		}
		image = append(image, HeroBannerImage)
	}

	result := database.DB.Create(&image)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create hero banner image: %w", result.Error)
	}

	var herobannerimage []common.HeroBannerImageResponse
	for _, image := range image {
		var heroimage = common.HeroBannerImageResponse{
			ImageURL: image.ImageURL,
		}
		herobannerimage = append(herobannerimage, heroimage)
	}
	return herobannerimage, nil
}

func (manager *heroBannerManager) UploadImageToS3(src io.ReadCloser, fileHeader *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(fileHeader.Filename)
	fileName := strings.ReplaceAll(fileHeader.Filename, ext, "")
	uniqueFileName := fmt.Sprintf("%s-%d%s", fileName, time.Now().UnixNano(), ext)
	dst, err := os.Create("./uploads" + uniqueFileName)
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
