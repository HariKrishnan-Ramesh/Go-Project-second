package managers

import (
	"errors"
	"fmt"
	"main/common"
	"main/database"
	"main/models"

	"gorm.io/gorm"
)

type CategoryManager interface {
	GetCategory() ([]common.CategoryResponse, error) 
}

type categoryManager struct {
	//dbclient
}

func NewCategotyManager() CategoryManager {
	return &categoryManager{}
}

func (manager *categoryManager) GetCategory() ([]common.CategoryResponse, error) {
	var categories []models.Category

	result := database.DB.Find(&categories)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []common.CategoryResponse{}, nil
		}
		return nil, fmt.Errorf("failed to retrieve categories %w", result.Error)
	}

	var categoryResponses []common.CategoryResponse
	for _, category := range categories {		
		categoryResponse := common.CategoryResponse{
			Id:            category.Id,
			CreatedAt:     category.CreatedAt.String(),
			UpdatedAt:     category.UpdatedAt.String(),
			CategoryName:  category.CategoryName,
			URLKey:        category.URLKey,
			Description:   category.Description,
			BannerImage:   category.BannerImage,
			CategoryImage: category.CategoryImage,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	return categoryResponses, nil
}
