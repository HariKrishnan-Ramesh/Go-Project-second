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
	GetHeroBanner() (common.HeroBannerResponse, error)
}

type heroBannerManager struct{}

func NewHeroBannerManager() HeroBannerManager {
	return &heroBannerManager{}
}

func (manager *heroBannerManager) GetHeroBanner() (common.HeroBannerResponse, error) {
	var heroBanner models.HeroBanner

	result := database.DB.Preload("Images").First(&heroBanner)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return common.HeroBannerResponse{}, fmt.Errorf("no hero banner found")
		}
		return common.HeroBannerResponse{}, fmt.Errorf("failed to retrieve hero banner: %w", result.Error)
	}

	heroBannerResponse := common.HeroBannerResponse{
		Id:          heroBanner.Id,
		Title:       heroBanner.Title,
		Description: heroBanner.Description,
		Images:      []common.HeroBannerImageResponse{},
	}

	for _, image := range heroBanner.Images {
		heroImageResponse := common.HeroBannerImageResponse{
			ImageURL: image.ImageURL,
		}
		heroBannerResponse.Images = append(heroBannerResponse.Images, heroImageResponse)
	}

	return heroBannerResponse, nil
}
