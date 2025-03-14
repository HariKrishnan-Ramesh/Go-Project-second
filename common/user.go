package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type requestResponse struct {
	Message string      `json:"message"`
	Status  uint        `json:"status"`
	Data    interface{} `json:"data,omitempty"`
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
