package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type UserCreationInput struct{
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null"`
	Password string `json:"not null"`
}


type requestResponse struct {
	Message string      `json:"message"`
	Status  uint        `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type UpdatePhoneInput struct {
	UserID uint `json:"userId" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type VerifyPhoneInput struct {
	UserID uint `json:"userid" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	OTP string `json:"otp" binding:"required"`
}

func SuccessResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		Message: msg,
		Status:  http.StatusOK,
	}

	ctx.JSON(http.StatusOK, response)
}

func BadResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		Message: msg,
		Status:  http.StatusOK,
	}

	ctx.JSON(http.StatusOK, response)
}

func SuccessResponseWithData(ctx *gin.Context, msg string, data interface{}) {
	response := requestResponse{
		Message: msg,
		Status:  http.StatusOK,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func InternalServerErrorResponse(ctx *gin.Context, msg string) {
	response := requestResponse{
		Message: msg,
		Status:  http.StatusInternalServerError,
	}

	ctx.JSON(http.StatusInternalServerError, response)
}


func NotFoundResponse(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, requestResponse{
		Status:  http.StatusNotFound,
		Message: message,
	})
}