package old_types

import (
	"time"

	"github.com/cloudlink-omega/storage/pkg/bitfield"
)

type Users struct {
	ID       string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Username string             `gorm:"unique;not null;type:tinytext;not null"`
	Email    string             `gorm:"unique;not null;type:tinytext;not null"`
	Password string             `gorm:"type:tinytext"`
	State    bitfield.Bitfield8 `gorm:"not null;default:0;type:tinyint(4);unsigned"`
	Created  int64              `gorm:"not null;type:bigint(20)"`
}

type Saves struct {
	UserID   string `gorm:"column:userid;foreignKey:UserID;type:char(26);not null;index:idx_saves_user_id"`
	GameID   string `gorm:"column:gameid;foreignKey:GameID;type:char(26);not null;index:idx_saves_game_id"`
	SlotID   uint8  `gorm:"column:slotid;not null;min:1;max:10;type:tinyint(4);unsigned"`
	Contents string `gorm:"not null;type:varchar(10000)"`

	User *Users `gorm:"constraint:OnDelete:CASCADE;"`
	Game *Games `gorm:"constraint:OnDelete:CASCADE;"`
}

type Games struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	DeveloperID string             `gorm:"column:developerid;foreignKey:DeveloperID;type:char(26);not null;index:idx_games_developer_id"`
	Name        string             `gorm:"type:tinytext;not null;default:'';index:idx_games_name"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;type:tinyint(4);unsigned"`
	Created     time.Time          `gorm:"not null"`

	Developer *Developers `gorm:"constraint:OnDelete:CASCADE;"`
}

type Developers struct {
	ID          string             `gorm:"primaryKey;type:char(26);unique;not null"`
	Name        string             `gorm:"type:tinytext;not null;default:''"`
	Description string             `gorm:"type:tinytext;not null;default:''"`
	State       bitfield.Bitfield8 `gorm:"not null;default:0;type:tinyint(4);unsigned"`
	Created     int64              `gorm:"not null;type:bigint(20)"`
}

type DeveloperMembers struct {
	DeveloperID string `gorm:"column:developerid;foreignKey:DeveloperID;type:char(26);not null;"`
	UserID      string `gorm:"column:userid;foreignKey:UserID;type:char(26);not null;"`

	Developer *Developers `gorm:"constraint:OnDelete:CASCADE;"`
	User      *Users      `gorm:"constraint:OnDelete:CASCADE;"`
}
