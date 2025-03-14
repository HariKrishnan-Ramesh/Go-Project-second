package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	Id                uint           `gorm:"primarykey" json:"id"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
	FirstName         string         `json:"firstName"`
	LastName          string         `json:"lastName"`
	Email             string         `json:"email"`
	Password          string         `json:"password"`
	Phone             string         `json:"phone"`
	VerificationToken string         `gorm:"column:verification_token" json:"-"`
	IsVerified        bool           `gorm:"column:is_verified;default:false" json:"isVerified"`
	Token             string         `gorm:"uniqueIndex:idx_users_token,length:191" json:"-"`
	Address           Address        `json:"address" gorm:"embedded"`
	Image             string         `json:"image,omitempty"`
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitempty"`
	City     string `json:"city"`
	District string `json:"district,omitempty"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Pin      string `json:"pin"`
}

type Otp struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"userId"`
	User      User      `gorm:"foreignKey:UserID"`
	OTP       string    `json:"otp"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}
