package handlers

import (
	"errors"
	"log"
	"main/common"
	"main/managers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	groupName  string
	otpManager managers.OtpManager
}

func NewOtpHandler(otpManager managers.OtpManager) *OtpHandler {
	return &OtpHandler{
		"api/otp",
		otpManager,
	}
}

func (otpHandler *OtpHandler) RegisterOtpApis(router *gin.Engine) {
	otpGroup := router.Group(otpHandler.groupName)
	otpGroup.POST("/send",otpHandler.SendOTP)
	otpGroup.POST("/verify",otpHandler.VerifyOTP)
}

func (otpHandler *OtpHandler) SendOTP(ctx *gin.Context) {
	userIDStr := ctx.Query("user_id")
	phoneNumber := ctx.Query("phone_number")

	if userIDStr == "" || phoneNumber == "" {
		common.BadResponse(ctx, "UserID and phone number are required")
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10,32)
	if err != nil {
		common.BadResponse(ctx," Invalid user ID")
		return
	}

	err = otpHandler.otpManager.SendOTP(uint(userID),phoneNumber)
	if err!=nil {
		log.Printf("Failed to send the otp: %v", err)
		common.InternalServerErrorResponse(ctx,"Failed to send OTP")
		return
	}

	common.SuccessResponse(ctx, "OTP Sent successfully")
}


func (otpHandler *OtpHandler) VerifyOTP(ctx *gin.Context) {
	phoneNumber := ctx.Query("phone_number")
	otp := ctx.Query("otp")

	if phoneNumber == "" || otp == "" {
		common.BadResponse(ctx, "Phone Number and OTP are required")
		return
	}

	err := otpHandler.otpManager.VerifyOTP(phoneNumber, otp)
	if err != nil {
		if errors.Is(err, managers.ErrInvalidOTP) {
			common.BadResponse(ctx, "Invalid OTP")
		} else if errors.Is(err, managers.ErrOTPExpired) {
			common.BadResponse(ctx, "OTP Expired")
		} else {
			log.Printf("Failed to verify OTP: %v", err)
			common.InternalServerErrorResponse(ctx," Failed to verify OTP")
		}
		return
	}
	common.SuccessResponse(ctx, "OTP verified successfully")
}