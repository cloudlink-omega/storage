package common

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) error {

	// Perform database migrations
	if err := db.AutoMigrate(
		&types.Event{},
		&types.User{},
		&types.UserEvent{},
		&types.Verification{},
		&types.RecoveryCode{},
		&types.UserGoogle{},
		&types.UserDiscord{},
		&types.UserGitHub{},
		&types.UserTOTP{},
		&types.UserSession{},
		&types.Developer{},
		&types.DeveloperEvent{},
		&types.DeveloperGame{},
		&types.DeveloperMember{},
		&types.DeveloperArt{},
		&types.UserGameSave{},
		&types.GameFeature{},
		&types.GameArt{},
	); err != nil {
		return err
	}

	// Seed the database with the test Developer and test Game IDs
	if err := db.FirstOrCreate(&types.Developer{
		Name: "Test Developer",
		ID:   "01HNPHQM5SPAG43J68R3NRX4M6",
	}).Error; err != nil {
		return err
	}

	if err := db.FirstOrCreate(&types.DeveloperGame{
		Name:        "Test Game",
		ID:          "01HNPHRWS0N0AYMM5K4HN31V4W",
		DeveloperID: "01HNPHQM5SPAG43J68R3NRX4M6",
		Description: "This is a sample game provided by the server for testing use.",
	}).Error; err != nil {
		return err
	}

	// Seed all events
	var events []map[string][]any = []map[string][]any{
		types.UserEvents,
		types.DeveloperEvents,
	}

	for _, entry := range events {
		for key, value := range entry {
			event := types.Event{
				ID:          key,
				Description: value[0].(string),
				LogLevel:    value[1].(uint8),
			}
			if err := db.FirstOrCreate(&event).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
