package managers

import (
	"errors"
	"fmt"
	"main/common"
	"main/database"
	"main/models"

	"gorm.io/gorm"
)

type HeroBannerManager interface {
	GetHeroBanner() ([]common.HeroBannerResponse, error)
	CreateHeroBanner(heroBanner models.HeroBanner) error
}

type heroBannerManager struct{}

func NewHeroBannerManager() HeroBannerManager {
	return &heroBannerManager{}
}

func (manager *heroBannerManager) CreateHeroBanner(heroBanner models.HeroBanner) error {
	result := database.DB.Create(&heroBanner)
	if result.Error != nil {
		return fmt.Errorf("failed to create hero banner: %w", result.Error)
	}
	return nil
}

func (manager *heroBannerManager) GetHeroBanner() ([]common.HeroBannerResponse, error) {
	var heroBanners []models.HeroBanner

	result := database.DB.Preload("Images").
		Where("is_active = ?", true).
		Order("position ASC").
		Find(&heroBanners)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []common.HeroBannerResponse{}, fmt.Errorf("no hero banner found")
		}
		return []common.HeroBannerResponse{}, fmt.Errorf("failed to retrieve hero banner: %w", result.Error)
	}

	if len(heroBanners) == 0 {
		return []common.HeroBannerResponse{}, fmt.Errorf("no active hero banners found")
	}

	var heroBannerResponses []common.HeroBannerResponse
	for _, heroBanner := range heroBanners {
		heroBannerResponse := common.HeroBannerResponse{
			Id:          heroBanner.Id,
			Title:       heroBanner.Title,
			Description: heroBanner.Description,
			ImageURL:    heroBanner.ImageURL,
			Position:    heroBanner.Position,
			Is_active:   heroBanner.Is_active,
		}
		heroBannerResponses = append(heroBannerResponses, heroBannerResponse)

	}
	return heroBannerResponses, nil

}


