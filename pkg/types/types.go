package types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type User struct {
	ID        string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Username  string             `gorm:"unique;not null;min:1;max:20"`
	Email     string             `gorm:"unique;not null;min:1;max:255"`
	Password  string             `gorm:"type:mediumtext"`
	Secret    string             `gorm:"type:mediumtext"`
	State     bitfield.Bitfield8 `gorm:"not null;default:0;"`
	AvatarID  *string
	BannerID  *string
	CreatedAt time.Time
	UpdatedAt time.Time

	Avatar        *Image          `gorm:"foreignKey:AvatarID;constraint:OnDelete:SET NULL;"`
	Banner        *Image          `gorm:"foreignKey:BannerID;constraint:OnDelete:SET NULL;"`
	UserGameSaves []*UserGameSave `gorm:"foreignKey:UserID"`
}

type UserSession struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"foreignKey:UserID;not null;"`
	UserAgent string `gorm:"mediumtext;not null"`
	Origin    string `gorm:"mediumtext;not null"`
	IP        string `gorm:"mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGoogle struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"foreignKey:UserID;not null;"`
	CreatedAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserDiscord struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"foreignKey:UserID;not null;"`
	CreatedAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGitHub struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"foreignKey:UserID;not null;"`
	CreatedAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserTOTP struct {
	UserID    string `gorm:"foreignKey:UserID;not null;"`
	Secret    string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Verification struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;"`
	Code      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type RecoveryCode struct {
	UserID    string `gorm:"foreignKey:UserID;type:char(26);not null;"`
	Code      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE;"`
}

type Achievement struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID          string `gorm:"foreignKey:UserID;not null;"`
	DeveloperGameID string `gorm:"foreignKey:DeveloperGameID;not null;"`
	Description     string `gorm:"type:tinytext;not null"`
	Points          uint64 `gorm:"not null;default:0"`
	IconID          *string
	CreatedAt       time.Time

	User          *User          `gorm:"constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"constraint:OnDelete:CASCADE;"`
	Icon          *Image         `gorm:"foreignKey:IconID;constraint:OnDelete:SET NULL;"`
}

type SystemEvent struct {
	ID         string `gorm:"primaryKey;type:char(26);unique;not null"`
	EventID    string `gorm:"foreignKey:EventID;"`
	Details    string `gorm:"type:tinytext"`
	Successful bool
	CreatedAt  time.Time

	Event *Event `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserEvent struct {
	ID         string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID     string `gorm:"foreignKey:UserID;not null;"`
	EventID    string `gorm:"foreignKey:EventID;"`
	Details    string `gorm:"type:tinytext"`
	Successful bool
	CreatedAt  time.Time

	User  *User  `gorm:"constraint:OnDelete:CASCADE;"`
	Event *Event `gorm:"constraint:OnDelete:CASCADE;"`
}

type DeveloperEvent struct {
	ID          string `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string `gorm:"foreignKey:DeveloperID;"`
	EventID     string `gorm:"foreignKey:EventID;"`
	Details     string `gorm:"type:tinytext"`
	Successful  bool
	CreatedAt   time.Time

	Developer *Developer `gorm:"constraint:OnDelete:CASCADE;"`
	Event     *Event     `gorm:"constraint:OnDelete:CASCADE;"`
}

type UserGameSave struct {
	UserID          string `gorm:"primaryKey;type:char(26);not null"`
	DeveloperGameID string `gorm:"primaryKey;type:char(26);not null"`
	SaveSlot        uint8  `gorm:"not null;min:1;max:10"`
	SaveData        string `gorm:"type:mediumtext"`

	CreatedAt time.Time
	UpdatedAt time.Time

	User          *User          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;constraint:OnDelete:CASCADE;"`
}

type Developer struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string             `gorm:"type:tinytext;not null;default:''"`
	Description string             `gorm:"type:mediumtext"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`
	BannerID    *string
	AvatarID    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Banner *Image `gorm:"foreignKey:BannerID;constraint:OnDelete:SET NULL;"`
	Avatar *Image `gorm:"foreignKey:AvatarID;constraint:OnDelete:SET NULL;"`
}

type DeveloperGame struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string             `gorm:"type:tinytext;not null;default:''"`
	Description string             `gorm:"type:mediumtext"`
	DeveloperID string             `gorm:"foreignKey:DeveloperID;"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`
	ThumbnailID *string
	CreatedAt   time.Time

	Thumbnail     *Image          `gorm:"foreignKey:ThumbnailID;constraint:OnDelete:SET NULL;"`
	Developer     *Developer      `gorm:"constraint:OnDelete:CASCADE;"`
	Features      []*FeatureTag   `gorm:"many2many:developer_game_features;"`
	UserGameSaves []*UserGameSave `gorm:"foreignKey:DeveloperGameID"`
}

type DeveloperMember struct {
	UserID      string             `gorm:"foreignKey:UserID;"`
	DeveloperID string             `gorm:"foreignKey:DeveloperID;"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`

	User      *User      `gorm:"constraint:OnDelete:CASCADE;"`
	Developer *Developer `gorm:"constraint:OnDelete:CASCADE;"`
}

type Image struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	Link      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Event struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null"`
	Description string `gorm:"type:tinytext"`
	LogLevel    uint8
}

type FeatureTag struct {
	ID          string `gorm:"not null;"`
	Description string `gorm:"type:tinytext"`
}
