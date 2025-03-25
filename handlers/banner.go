package handlers

import (
	"log"
	"main/common"
	"main/managers"

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
	heroBannerGroup.GET("", handler.GetHeroBanner) // Modified to be a GET without a parameter. Remove the Id from route.
}



// GetHeroBanner retrieves the hero banner. No ID is needed. Returns the single banner.
func (handler *HeroBannerHandler) GetHeroBanner(ctx *gin.Context) {
	heroBanner, err := handler.heroBannerManager.GetHeroBanner()
	if err != nil {
		log.Printf("Error retrieving hero banner: %v", err)
		if err.Error() == "no hero banner found"{
			common.NotFoundResponse(ctx,"No hero banner found")
			return
		}
		common.InternalServerErrorResponse(ctx, "Failed to retrieve hero banner")
		return
	}

	common.SuccessResponseWithData(ctx, "Hero banner retrieved successfully", heroBanner)
}