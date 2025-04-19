package types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type Developer struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string             `gorm:"type:varchar(255);not null;default:''"`
	Description string             `gorm:"type:varchar(255);not null;default:''"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Games   []*DeveloperGame
	Members []*DeveloperMember
	Art     []*DeveloperArt
}

type DeveloperGame struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string             `gorm:"foreignKey:DeveloperID;type:char(26);not null;index:idx_developer_game_developer_id"`
	Name        string             `gorm:"type:varchar(255);not null;default:'';index:idx_developer_game_name"`
	Description string             `gorm:"type:varchar(255);not null;default:''"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Developer   Developer `gorm:"constraint:OnDelete:CASCADE;"`
	GameSave    []*UserGameSave
	GameFeature []*GameFeature
	GameArt     []*GameArt
}

type DeveloperMember struct {
	UserID      User                `gorm:"foreignKey:UserID;type:char(26);not null;index:idx_developer_member_user_id"`
	DeveloperID string              `gorm:"foreignKey:DeveloperID;type:char(26);not null;index:idx_developer_member_developer_id"`
	State       bitfield.Bitfield16 `gorm:"not null;default:0"`
	Linked      time.Time           `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Developer Developer `gorm:"constraint:OnDelete:CASCADE;"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;"`
}

type GameFeature struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperGameID string `gorm:"foreignKey:DeveloperGameID;type:char(26);not null;index:idx_game_feature_developer_game_id"`
	TagID           string `gorm:"type:varchar(20);unique;not null;default:''"`
	TagDescription  string `gorm:"type:varchar(255);not null;default:''"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	DeveloperGame DeveloperGame `gorm:"constraint:OnDelete:CASCADE;"`
}

type GameArt struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperGameID string `gorm:"foreignKey:DeveloperGameID;type:char(26);not null;index:idx_game_art_developer_game_id"`
	ArtType         string `gorm:"type:varchar(20);not null;default:''"`
	ArtLink         string `gorm:"type:varchar(255);not null;default:''"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	DeveloperGame DeveloperGame `gorm:"constraint:OnDelete:CASCADE;"`
}

type DeveloperArt struct {
	ID          string `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string `gorm:"foreignKey:DeveloperID;type:char(26);not null;index:idx_developer_art_developer_id"`
	ArtType     string `gorm:"type:varchar(20);not null;default:''"`
	ArtLink     string `gorm:"type:varchar(255);not null;default:''"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Developer Developer `gorm:"constraint:OnDelete:CASCADE;"`
}

/* {
	"ImageURL":   "",
	"Title":      ". . .",
	"Text":       "",
	"Creator":    "",
	"FooterText": "",
	"Enabled":    false,
	"IsNew":      false,
	"ID":         "",
	"Features":   []map[string]any{},
} */
