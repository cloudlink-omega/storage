package types

import (
	"time"
)

type UserGameSave struct {
	UserID          string `gorm:"type:char(26);not null;foreignKey:UserID"`
	DeveloperGameID string `gorm:"type:char(26);not null;foreignKey:DeveloperGameID"`
	SaveSlot        uint8  `gorm:"not null,min=1,max=10"`
	SaveData        string `gorm:"not null,max=10000"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	User          User          `gorm:"constraint:OnDelete:CASCADE;"`
	DeveloperGame DeveloperGame `gorm:"constraint:OnDelete:CASCADE;"`
}
