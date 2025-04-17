package types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type Developer struct {
	ID      string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name    string             `gorm:"type:varchar(255);not null"`
	Created *time.Time         `gorm:"not null"`
	State   bitfield.Bitfield8 `gorm:"not null;default:0"`

	Games []*Game `gorm:"foreignKey:DeveloperID"`
}

type Game struct {
	ID      string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name    string             `gorm:"type:varchar(255);not null"`
	Created *time.Time         `gorm:"not null"`
	State   bitfield.Bitfield8 `gorm:"not null;default:0"`
}
