package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	FirstName string         `json:"firstname"`
	LastName  string         `json:"lastname"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Phone     string         `json:"phone"`
	IsSeller  *bool          `json:"isSeller" gorm:"default:null"`
	IsBuyer   *bool          `json:"isBuyer"  gorm:"default:null"`
	Address   Address        `json:"address" gorm:"embedded"`
	Image     string         `json:"image,omitempty"`
	IsAdmin   *bool          `json:"isadmin" gorm:"default:null"`
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

type Configuration struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	LogoURL   string         `json:"logoURL"`
}

type HeroBanner struct {
	Id          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	ImageURL    string         `json:"imageURL"`
	Position    uint           `json:"position"`
	Is_active   bool           `json:"is_active"`
}


type Category struct {
	Id            uint           `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	CategoryName  string         `json:"category_name"`
	URLKey        string         `json:"url_key"`
	Description   string         `json:"description"`
	BannerImage   []string       `gorm:"type:json" json:"banner_image"`
	CategoryImage []string       `gorm:"type:json" json:"category_image"`
}


// type Wishlist struct {
// 	Id        uint      `gorm:"primaryKey" json:"id"`
// 	CreatedAt time.Time `json:"createdAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// 	UserID    uint      `json:"userID"`
// 	PropertyID 
// 	User      User      `json:"user" gorm:"foreignKey:UserID"`
// }