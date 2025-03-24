package handlers

import (
	"fmt"
	"log"
	"main/common"
	"main/managers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	groupName    string
	adminManager managers.AdminManager
}

func NewAdminHandler(adminManager managers.AdminManager) *AdminHandler {
	return &AdminHandler{
		"api/admin",
		adminManager,
	}
}


func (handler *AdminHandler) RegisterAdminApis(router *gin.Engine){
	adminGroup := router.Group(handler.groupName)
	adminGroup.Use(handler.AdminAuthMiddleware())
	adminGroup.POST("/logo",handler.UploadLogo)
	adminGroup.GET("/logo",handler.GetLogo)
}

func (handler *AdminHandler) AdminAuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("userID")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized"})
			return
		}

		if userID == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized"})
			return
		}
		isAdmin, err := handler.adminManager.IsAdmin(uint(userID.(int)))
		if err != nil{
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"There is a problem in User!. Please try again Later"})
			return
		}
		if !isAdmin{
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error":"Forbidden"})
			return
		}

		ctx.Next()
	}
}


func (handler *AdminHandler) UploadLogo(ctx *gin.Context) {
	file, err := ctx.FormFile("logo")
	if err != nil {
		common.BadResponse(ctx, "Failed to get logo from request")
		return
	}

	src, err := file.Open()
	if err != nil {
		common.InternalServerErrorResponse(ctx, "Failed to open logo file")
		return
	}
	defer src.Close()

	response, err := handler.adminManager.UploadLogo(src, file)
	if err != nil {
		// Print error details to logs
		log.Printf("UploadLogo Error: %v", err) 
		common.InternalServerErrorResponse(ctx, fmt.Sprintf("Failed to upload logo: %v", err))
		return
	}
	common.SuccessResponseWithData(ctx, "Logo uploaded successfully", response)
}

func (handler *AdminHandler) GetLogo(ctx *gin.Context) {
	response, err := handler.adminManager.GetLogo()
	if err != nil {
		common.BadResponse(ctx, "No logo found. Please upload a logo.")
		return
	}
	common.SuccessResponseWithData(ctx, "Logo retrieved successfully", response)
}