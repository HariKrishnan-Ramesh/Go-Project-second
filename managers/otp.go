package managers

import (
	"errors"
	"fmt"
	"log"
	"main/common"
	"main/database"
	"main/models"
	"os"
	"time"

	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)


var (
	ErrInvalidOTP = errors.New("invalid OTP")
	ErrOTPExpired = errors.New("OTP Expired")
)

type OtpManager interface{
	SendOTP(userID uint, phoneNumber string) error
	VerifyOTP(phoneNumber string, otp string) error
}

type otpManager struct{

}

func NewOtpManager() OtpManager {
	return &otpManager{}
}

const otpLength = 6 
const otpExpiration = 5 * time.Minute

func (otpManager *otpManager) SendOTP(userID uint, phoneNumber string) error {

	otp := common.GenerateOTP(otpLength)
	expiresAt := time.Now().Add(otpExpiration)

	otpRecord := models.Otp{
		UserID : userID,
		OTP : otp,
		CreatedAt : time.Now(),
		ExpiresAt: expiresAt,
	}

	result := database.DB.Create(&otpRecord)
	if result.Error != nil {
		return fmt.Errorf("failed to create Otp Record: %w",result.Error)
	}

	err := sendOTP(phoneNumber, otp)
	if err != nil {
		deleteErr := database.DB.Delete(&otpRecord).Error
		if deleteErr != nil {
			log.Printf("Failed to delete OTP record after sending failed: %v", deleteErr)

		}
		return fmt.Errorf("failed to send OTP via Twilio: %w", err)
	}
	return nil
}

func (otpManager *otpManager) VerifyOTP(phoneNumber string, otp string) error {
	return nil
}

func sendOTP(phoneNumber, otp string) error {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	messageInput := &openapi.CreateMessageParams{}
	messageInput.SetTo(phoneNumber)
	messageInput.SetFrom(twilioPhoneNumber)
	messageInput.SetBody(fmt.Sprintf("Your OTP is: %s", otp))

	_, err := client.Api.CreateMessage(messageInput)
	if err != nil {
		log.Printf("twilio error %s", err.Error())
		return err
	}

	log.Println("OTP sent successfully")

	return nil
}
