package handlers

import (
	"log"
	"main/common"
	"main/managers"
	"net/http"
	"strconv"

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
	heroBannerGroup.POST("",handler.CreateHeroBanner)
}


func (handler *HeroBannerHandler) CreateHeroBanner(ctx *gin.Context) {
	var input common.HeroBannerCreationInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	heroBanner, err := handler.heroBannerManager.CreateHeroBanner(input)
	if err != nil{
		log.Printf("Error creating hero Banner: %v", err)
		common.InternalServerErrorResponse(ctx, "Failed to create hero banner")
		return
	}

	common.SuccessResponseWithData(ctx, "Hero banner created successfully", heroBanner)
}

func (handler *HeroBannerHandler) GetHeroBanner(ctx *gin.Context) {
	heroBannerIdStr := ctx.Param("herobannerId")
	heroBannerId, err := strconv.Atoi(heroBannerIdStr)
	if err != nil {
		common.BadResponse(ctx, "Invalid Hero Banner ID")
		return
	}

	heroBanner, err := handler.heroBannerManager.GetHeroBanner(uint(heroBannerId))
	if err != nil {
		log.Printf("Error retrieving hero banner: %v", err)
		common.InternalServerErrorResponse(ctx, "Failed to retrieve hero banner")
		return
	}

	common.SuccessResponseWithData(ctx, "Hero banner retrieved successfully", heroBanner)
}

func (handler *HeroBannerHandler) UploadImages(ctx *gin.Context) {
	heroBannerIdStr := ctx.Param("herobannerId")
	heroBannerId, err := strconv.Atoi(heroBannerIdStr)
	if err != nil {
		common.BadResponse(ctx, "Invalid Hero Banner ID")
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusBadRequest,"get form err: %s",err.Error())
		return 
	}
	files := form.File["images"]

	var uploadedImageURLs []string

	for _, file := range files {
		imageURL, err := handler.UploadImageToS3(ctx,file)

		if err != nil {
			log.Printf("Error uploading image: %v", err)
			continue 
		}
		uploadedImageURLs = append(uploadedImageURLs, imageURL)
	}
	heroBannerImages, err := handler.heroBannerManager.AddImages(uint(heroBannerId),uploadedImageURLs)
	if err !=nil {
		log.Printf("Error creating property images in DB: %v", err)
		common.InternalServerErrorResponse(ctx, "Failed to create images in DB")
		return
	}

	common.SuccessResponseWithData(ctx, "Image successfully uploaded",heroBannerImages)
}

