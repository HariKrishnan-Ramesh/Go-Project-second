package handlers

import (
	"log"
	"main/common"
	"main/managers"
	"main/models"

	"github.com/gin-gonic/gin"
)

type HeroBannerHandler struct {
	groupName         string
	heroBannerManager managers.HeroBannerManager
}

func NewHeroBannerHandler(heroBannerManager managers.HeroBannerManager) *HeroBannerHandler {
	return &HeroBannerHandler{
		"api/herobanner",
		heroBannerManager,
	}
}

func (handler *HeroBannerHandler) RegisterHeroBannerApis(router *gin.Engine) {
	heroBannerGroup := router.Group(handler.groupName)
	heroBannerGroup.GET("", handler.GetHeroBanner)
	heroBannerGroup.POST("", handler.CreateHeroBanner)
}

func (handler *HeroBannerHandler) CreateHeroBanner(ctx *gin.Context) {
	var newHeroBanner models.HeroBanner

	if err := ctx.ShouldBindJSON(&newHeroBanner); err != nil {
		log.Printf("Error binding JSON: %v", err)
		common.BadResponse(ctx, "Invalid request body")
		return
	}

	err := handler.heroBannerManager.CreateHeroBanner(newHeroBanner)
	if err != nil {
		log.Printf("Error creating hero banner: %v", err)
		common.InternalServerErrorResponse(ctx, "Failed to create hero banner")
		return
	}

	common.SuccessResponse(ctx, "Hero banner created successfully")
}

func (handler *HeroBannerHandler) GetHeroBanner(ctx *gin.Context) {
	heroBanner, err := handler.heroBannerManager.GetHeroBanner()
	if err != nil {
		log.Printf("Error retrieving hero banner: %v", err)
		if err.Error() == "no hero banner found" {
			common.NotFoundResponse(ctx, "No hero banner found")
			return
		}
		common.InternalServerErrorResponse(ctx, "Failed to retrieve hero banner")
		return
	}

	common.SuccessResponseWithData(ctx, "Hero banner retrieved successfully", heroBanner)
}
