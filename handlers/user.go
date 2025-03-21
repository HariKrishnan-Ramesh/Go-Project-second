package handlers

import (
	"log"
	"main/common"
	"main/database"
	"main/managers"
	"main/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


type UserHandler struct{
	groupName string
	UserManager managers.UserManager
}

func NewUserHandlerForm(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/user",
		userManager,
	}
}


func (userHandler *UserHandler) RegisterUserApis(router *gin.Engine) {
	userGroup := router.Group(userHandler.groupName)
	userGroup.POST("/signup",userHandler.SignUp)
}

func (userHandler *UserHandler) SignUp(ctx *gin.Context) {
	var input common.UserCreationInput

	if err:= ctx.ShouldBindJSON(&input) ; err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		common.InternalServerErrorResponse(ctx,"Failed to hash password")
		return
	}

	user := models.User{
		FirstName:input.Name,
		Email: input.Email,
		Password: string(hashedPassword),
		Phone: input.Phone,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		common.InternalServerErrorResponse(ctx, "Failed to create user")
		return
	}

	common.SuccessResponseWithData(ctx, "User created Successfully", user)
}


func (userHandler *UserHandler) UpdatePhoneNumber(ctx *gin.Context) {
	var input common.UpdatePhoneInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	otpManager := managers.NewOtpManager()
	err := otpManager.SendOTP(input.UserID, input.PhoneNumber)
	if err != nil {
		log.Printf("Failed to send OTP: %v" , err)
		common.InternalServerErrorResponse(ctx, "Failed to send OTP")
		return
	}

	common.SuccessResponse(ctx, "OTP sent to the new phone number. Verify to update")
}


func (userHandler *UserHandler) VerifyPhoneNumber(ctx *gin.Context) {
	var input common.VerifyPhoneInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.BadResponse(ctx, err.Error())
		return
	}

	otpManager := managers.NewOtpManager()
	err := otpManager.VerifyOTP(input.PhoneNumber, input.OTP)
	if err != nil {
		common.BadResponse(ctx, "Invalid or expired OTP")
		return
	}

	
}