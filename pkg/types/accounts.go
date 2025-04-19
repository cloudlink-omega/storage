package types

import (
	"fmt"
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type User struct {
	ID        string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Username  string             `gorm:"unique;not null;type:varchar(30)"`
	Email     string             `gorm:"unique;not null;type:varchar(255)"`
	Password  string             `gorm:"type:varchar(255)"`
	State     bitfield.Bitfield8 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Google       *UserGoogle
	Discord      *UserDiscord
	GitHub       *UserGitHub
	TOTP         *UserTOTP
	Verify       *Verification
	Recovery     []*RecoveryCode
	Developer    *DeveloperMember
	UserGameSave []*UserGameSave
}

type UserGoogle struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);unique;not null"`
	GoogleID  string `gorm:"primaryKey;type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserDiscord struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);unique;not null"`
	DiscordID string `gorm:"primaryKey;type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGitHub struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);unique;not null"`
	GitHubID  string `gorm:"primaryKey;type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserTOTP struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);unique;not null"`
	Secret    string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Verification struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null"`
	Code      string `gorm:"type:char(6);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type RecoveryCode struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null"`
	Code      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *User) String() string {
	return fmt.Sprintf("[ID: %s, Username: %s, Password: %s, Email: %s, State: %d]", u.ID, u.Username, u.Password, u.Email, u.State)
}
