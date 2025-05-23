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

	Avatar          *Image             `gorm:"foreignKey:AvatarID;references:ID;constraint:OnDelete:SET NULL;"`
	Banner          *Image             `gorm:"foreignKey:BannerID;references:ID;constraint:OnDelete:SET NULL;"`
	UserGameSaves   []*UserGameSave    `gorm:"foreignKey:UserID"`
	DeveloperMember []*DeveloperMember `gorm:"foreignKey:UserID"`
	GameComments    []*GameComment     `gorm:"foreignKey:UserID"`
}

type UserSession struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID    string `gorm:"not null"`
	UserAgent string `gorm:"mediumtext;not null"`
	Origin    string `gorm:"mediumtext;not null"`
	IP        string `gorm:"mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserGoogle struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserDiscord struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserGitHub struct {
	ID        string `gorm:"primaryKey;type:varchar(255);not null;unique;"`
	UserID    string `gorm:"not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type UserTOTP struct {
	UserID    string `gorm:"not null"`
	Secret    string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Verification struct {
	UserID    string `gorm:"type:char(26);not null"`
	Code      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type RecoveryCode struct {
	UserID    string `gorm:"type:char(26);not null"`
	Code      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
}

type Achievement struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID          string `gorm:"not null"`
	DeveloperGameID string `gorm:"not null"`
	Description     string `gorm:"type:tinytext;not null"`
	Points          uint64 `gorm:"not null;default:0"`
	IconID          *string
	CreatedAt       time.Time

	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
	Icon          *Image         `gorm:"foreignKey:IconID;references:ID;constraint:OnDelete:SET NULL;"`
}

type SystemEvent struct {
	ID         string `gorm:"primaryKey;type:char(26);unique;not null"`
	EventID    string
	Details    string `gorm:"type:tinytext"`
	Successful bool
	CreatedAt  time.Time

	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserEvent is used to log changes to a user account, ranging from authentication and account changes to account errors.
type UserEvent struct {
	ID         string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID     string `gorm:"not null"`
	EventID    string
	Details    string `gorm:"type:tinytext"`
	Successful bool
	CreatedAt  time.Time

	User  *User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Event *Event `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserReport represents a user-submitted report on another user.
// By specifying a ReportTag, the user can specify the kind of the report.
// If the ReportTag is null, the user can fill out the ReportDetails field for a custom report.
type UserReport struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	UserID          string `gorm:"type:char(26);not null"`
	SubmittedUserID string `gorm:"type:char(26);not null"`
	ReportTagID     *string
	Details         string `gorm:"type:mediumtext"`
	CreatedAt       time.Time

	SubmittedUser *User      `gorm:"foreignKey:SubmittedUserID;references:ID;constraint:OnDelete:CASCADE;"`
	User          *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	ReportTag     *ReportTag `gorm:"foreignKey:ReportTagID;references:ID;constraint:OnDelete:CASCADE;"`
}

// DeveloperGameReport represents a user-submitted report on a developer's game.
// By specifying a ReportTag, the user can specify the kind of the report.
// If the ReportTag is null, the user can fill out the ReportDetails field for a custom report.
type DeveloperGameReport struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	SubmittedUserID string `gorm:"not null"`
	DeveloperGameID string `gorm:"not null"`
	ReportTagID     *string
	Details         string `gorm:"type:mediumtext"`
	CreatedAt       time.Time

	SubmittedUser *User          `gorm:"foreignKey:SubmittedUserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
	ReportTag     *ReportTag     `gorm:"foreignKey:ReportTagID;references:ID;constraint:OnDelete:CASCADE;"`
}

// DeveloperReport represents a user-submitted report on a developer's profile.
// By specifying a ReportTag, the user can specify the kind of the report.
// If the ReportTag is null, the user can fill out the ReportDetails field for a custom report.
type DeveloperReport struct {
	ID              string `gorm:"primaryKey;type:char(26);unique;not null"`
	SubmittedUserID string `gorm:"not null"`
	DeveloperID     string `gorm:"not null"`
	ReportTagID     *string
	Details         string `gorm:"type:mediumtext"`
	CreatedAt       time.Time

	SubmittedUser *User      `gorm:"foreignKey:SubmittedUserID;references:ID;constraint:OnDelete:CASCADE;"`
	Developer     *Developer `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
	ReportTag     *ReportTag `gorm:"foreignKey:ReportTagID;references:ID;constraint:OnDelete:CASCADE;"`
}

// DeveloperEvent is used to log changes to a developer profile, ranging from memberships to approvals.
type DeveloperEvent struct {
	ID          string `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string
	EventID     string
	Details     string `gorm:"type:tinytext"`
	Successful  bool
	CreatedAt   time.Time

	Developer *Developer `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
	Event     *Event     `gorm:"foreignKey:EventID;references:ID;constraint:OnDelete:CASCADE;"`
}

// UserGameSave is used to store a user's game save data.
type UserGameSave struct {
	UserID          string `gorm:"primaryKey;type:char(26);not null"`
	DeveloperGameID string `gorm:"primaryKey;type:char(26);not null"`
	SaveSlot        uint8  `gorm:"not null;min:1;max:10"`
	SaveData        string `gorm:"type:mediumtext"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Developer represents a game developer.
type Developer struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string             `gorm:"type:tinytext;not null;default:''"`
	Description string             `gorm:"type:mediumtext"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`
	BannerID    *string
	AvatarID    *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Banner           *Image             `gorm:"foreignKey:BannerID;references:ID;constraint:OnDelete:SET NULL;"`
	Avatar           *Image             `gorm:"foreignKey:AvatarID;references:ID;constraint:OnDelete:SET NULL;"`
	DeveloperMembers []*DeveloperMember `gorm:"foreignKey:DeveloperID"`
}

// DeveloperGame represents a game created by a developer.
type DeveloperGame struct {
	ID          string `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string `gorm:"type:tinytext;not null;default:''"`
	Description string `gorm:"type:mediumtext"`
	DeveloperID string
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`
	ThumbnailID *string
	CreatedAt   time.Time

	Thumbnail     *Image          `gorm:"foreignKey:ThumbnailID;references:ID;constraint:OnDelete:SET NULL;"`
	Developer     *Developer      `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
	Features      []*FeatureTag   `gorm:"many2many:developer_game_features;"`
	UserGameSaves []*UserGameSave `gorm:"foreignKey:DeveloperGameID"`
	GameComments  []*GameComment  `gorm:"foreignKey:DeveloperGameID"`
}

// GameComment represents a comment on a game by a user.
type GameComment struct {
	ID              string `gorm:"primaryKey;type:char(26);not null"`
	UserID          string `gorm:"type:char(26);not null"`
	DeveloperGameID string `gorm:"type:char(26);not null"`
	ParentID        *string
	Content         string `gorm:"type:mediumtext;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Relationships
	User          *User          `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	DeveloperGame *DeveloperGame `gorm:"foreignKey:DeveloperGameID;references:ID;constraint:OnDelete:CASCADE;"`
	Parent        *GameComment   `gorm:"foreignKey:ParentID;references:ID;constraint:OnDelete:CASCADE;"`
	Replies       []*GameComment `gorm:"foreignKey:ParentID"`
}

// DeveloperMember represents memberships between a user and a developer account.
type DeveloperMember struct {
	UserID      string             `gorm:"not null"`
	DeveloperID string             `gorm:"not null"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;"`

	User      *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	Developer *Developer `gorm:"foreignKey:DeveloperID;references:ID;constraint:OnDelete:CASCADE;"`
}

// Image represents an image stored on the server's hosted folder.
type Image struct {
	ID        string `gorm:"primaryKey;type:char(26);unique;not null"`
	Link      string `gorm:"type:mediumtext;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Event is a generic entity used to de-duplicate events across the system.
type Event struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null"`
	Description string `gorm:"type:tinytext"`
	LogLevel    uint8
}

// FeatureTag is a generic entity used to de-duplicate features for games.
type FeatureTag struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null;"`
	Description string `gorm:"type:tinytext"`
}

// ReportTag is a generic entity used to de-duplicate report types.
type ReportTag struct {
	ID          string `gorm:"primaryKey;type:varchar(50);unique;not null;"`
	Description string `gorm:"type:tinytext"`
	IsUser      bool
	IsDeveloper bool
	IsGame      bool
}
