package types

import (
	"fmt"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type User struct {
	ID       string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Username string             `gorm:"type:varchar(30)"`
	Password string             `gorm:"type:varchar(255)"`
	Email    string             `gorm:"type:varchar(255)"`
	State    bitfield.Bitfield8 `gorm:"not null;default:0"`

	Google   *UserGoogle     `gorm:"foreignKey:UserID"`
	Discord  *UserDiscord    `gorm:"foreignKey:UserID"`
	GitHub   *UserGitHub     `gorm:"foreignKey:UserID"`
	TOTP     *UserTOTP       `gorm:"foreignKey:UserID"`
	Verify   *Verification   `gorm:"foreignKey:UserID"`
	Recovery []*RecoveryCode `gorm:"foreignKey:UserID"`
}

type UserGoogle struct {
	UserID string `gorm:"primaryKey;type:char(26);unique;not null"`
	ID     string `gorm:"type:varchar(255);unique;not null"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserDiscord struct {
	UserID string `gorm:"primaryKey;type:char(26);unique;not null"`
	ID     string `gorm:"type:varchar(255);unique;not null"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGitHub struct {
	UserID string `gorm:"primaryKey;type:char(26);unique;not null"`
	ID     string `gorm:"type:varchar(255);unique;not null"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserTOTP struct {
	UserID string `gorm:"primaryKey;type:char(26);unique;not null"`
	Secret string `gorm:"type:varchar(255);not null"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Verification struct {
	UserID string `gorm:"primaryKey;type:char(26);unique;not null"`
	Code   string `gorm:"type:char(6);not null"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type RecoveryCode struct {
	UserID string `gorm:"type:char(26);not null"`
	Code   string `gorm:"type:varchar(50);not null"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *User) String() string {
	return fmt.Sprintf("[ID: %s, Username: %s, Password: %s, Email: %s, State: %d]", u.ID, u.Username, u.Password, u.Email, u.State)
}
