package handlers

import (
	"log"
	"main/common"
	"main/managers"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	groupName    string
	categoryManager managers.CategoryManager
}

func NewCategoryHandler(categoryManager managers.CategoryManager) *CategoryHandler {
	return &CategoryHandler{
		"api/category",
		categoryManager,
	}
}


func (handler *CategoryHandler) RegisterCategoryApis(router *gin.Engine) {
	categoryGroup := router.Group(handler.groupName)
	categoryGroup.GET("",handler.GetCategory)
}


func (handler *CategoryHandler) GetCategory(ctx *gin.Context) {
	categories, err := handler.categoryManager.GetCategory()
	if err != nil {
		log.Printf("Error retrieving categories: %v", err)
		common.InternalServerErrorResponse(ctx,"Failed to retrieve categories")
		return
	}

	common.SuccessResponseWithData(ctx, "Categories retrieved successfully", categories)
}