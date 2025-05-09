package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"` // "-" means this field won't appear in JSON
	Role      string    `gorm:"not null;default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	VerificationToken string  `gorm:"index" json:"-"`
	TokenExpiresAt time.Time `json:"-"`
	ResetPasswordToken string `gorm:"index" json:"-"`
	ResetTokenExpiresAt time.Time `json:"-"`
	RefreshToken     string    `gorm:"size:512" json:"-"`
	RefreshExpiresAt time.Time `json:"-"`
}

