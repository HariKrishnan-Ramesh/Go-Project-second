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
	GetHeroBanner() (common.HeroBannerResponse, error) // Modified to fetch the single existing banner
}

type heroBannerManager struct{}

func NewHeroBannerManager() HeroBannerManager {
	return &heroBannerManager{}
}

// GetHeroBanner retrieves the single existing hero banner.  It assumes there is only ONE.  Error handling is crucial here.
func (manager *heroBannerManager) GetHeroBanner() (common.HeroBannerResponse, error) {
	var heroBanner models.HeroBanner

	// Assuming you only have ONE HeroBanner record in the database.
	result := database.DB.Preload("Images").First(&heroBanner)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Handle the case where NO HeroBanner exists.  This is important!
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