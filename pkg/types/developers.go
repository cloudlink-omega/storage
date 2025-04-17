package types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type Developer struct {
	ID      string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name    string             `gorm:"type:varchar(255);not null"`
	Created time.Time          `gorm:"not null"`
	State   bitfield.Bitfield8 `gorm:"not null;default:0"`

	Games   []*DeveloperGame   `gorm:"foreignKey:DeveloperID"`
	Members []*DeveloperMember `gorm:"foreignKey:DeveloperID"`
}

type DeveloperGame struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string             `gorm:"type:char(26);not null"`
	Name        string             `gorm:"type:varchar(255);not null"`
	Created     time.Time          `gorm:"not null"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0"`

	Developer Developer `gorm:"constraint:OnDelete:CASCADE;"`
}

type DeveloperMember struct {
	UserID      User                `gorm:"type:char(26);not null"`
	DeveloperID string              `gorm:"type:char(26);not null"`
	State       bitfield.Bitfield16 `gorm:"not null;default:0"`
	Linked      time.Time           `gorm:"not null"`

	Developer Developer `gorm:"constraint:OnDelete:CASCADE;"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;"`
}
