package types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type Developer struct {
	ID        string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name      string             `gorm:"type:varchar(255);not null"`
	State     bitfield.Bitfield8 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Games   []*DeveloperGame
	Members []*DeveloperMember
}

type DeveloperGame struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string             `gorm:"foreignKey:DeveloperID;type:char(26);not null"`
	Name        string             `gorm:"type:varchar(255);not null"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Developer Developer
	GameSave  []*UserGameSave
}

type DeveloperMember struct {
	UserID      User                `gorm:"foreignKey:UserID;type:char(26);not null"`
	DeveloperID string              `gorm:"foreignKey:DeveloperID;type:char(26);not null"`
	State       bitfield.Bitfield16 `gorm:"not null;default:0"`
	Linked      time.Time           `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Developer Developer `gorm:"constraint:OnDelete:CASCADE;"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;"`
}
