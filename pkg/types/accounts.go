package types

import (
	"fmt"
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type User struct {
	ID       string `gorm:"primaryKey;type:char(26);unique;not null"`
	Username string `gorm:"unique;not null;type:varchar(30)"`
	Email    string `gorm:"unique;not null;type:varchar(255)"`

	// A scrypted hash of the user's password
	Password string `gorm:"type:varchar(255)"`

	// A base64 encoded secret based on a randomly generated 256-bit key
	// that is encrypted using the server's secret key.
	Secret string `gorm:"type:varchar(255)"`

	// A 8-bit field of flags that are used to determine the properties of the user.
	State bitfield.Bitfield8 `gorm:"not null;default:0"`

	// The date and time the user was created and last updated
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
	UserSession  []*UserSession
}

type UserSession struct {
	ID        string `gorm:"primaryKey;type:varchar(26);unique;not null"`
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_user_session_user_id"`
	UserAgent string `gorm:"type:varchar(255);not null"`
	Origin    string `gorm:"type:varchar(255);not null"`
	IP        string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGoogle struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_user_google_user_id"`
	GoogleID  string `gorm:"primaryKey;type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserDiscord struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_user_discord_user_id"`
	DiscordID string `gorm:"primaryKey;type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGitHub struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_user_git_hub_user_id"`
	GitHubID  string `gorm:"primaryKey;type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserTOTP struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_user_totp_user_id"`
	Secret    string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Verification struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_verification_user_id"`
	Code      string `gorm:"type:char(6);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

type RecoveryCode struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_recovery_code_user_id"`
	Code      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}

func (u *User) String() string {
	return fmt.Sprintf("[ID: %s, Username: %s, Email: %s, State: %d]", u.ID, u.Username, u.Email, u.State)
}
