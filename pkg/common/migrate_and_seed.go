package common

import (
	"time"

	"github.com/cloudlink-omega/accounts/pkg/database"
	"github.com/cloudlink-omega/storage/pkg/old_types"
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func MigrateAndSeed(db *gorm.DB) error {

	if db == nil {
		panic("Got nil database")
	}

	// Perform database migrations
	if err := db.AutoMigrate(
		&types.Event{},
		&types.ReportTag{},
		&types.SystemEvent{},
		&types.User{},
		&types.Verification{},
		&types.RecoveryCode{},
		&types.UserGoogle{},
		&types.UserDiscord{},
		&types.UserGitHub{},
		&types.UserTOTP{},
		&types.UserSession{},
		&types.UserEvent{},
		&types.UserReport{},
		&types.Developer{},
		&types.DeveloperEvent{},
		&types.DeveloperGame{},
		&types.DeveloperMember{},
		&types.DeveloperGameReport{},
		&types.DeveloperReport{},
		&types.GameComment{},
		&types.Achievement{},
		&types.UserGameSave{},
		&types.Image{},
		&types.FeatureTag{},
	); err != nil {
		return err
	}

	// Seed all feature tags
	for key, entry := range types.GameFeatureTags {
		tag := &types.FeatureTag{
			ID:          key,
			Description: entry,
		}
		if err := db.FirstOrCreate(&tag).Error; err != nil {
			return err
		}
	}

	// Seed all events
	var events []map[string][]any = []map[string][]any{
		types.UserEvents,
		types.DeveloperEvents,
	}

	for _, entry := range events {
		for key, value := range entry {
			event := &types.Event{
				ID:          key,
				Description: value[0].(string),
				LogLevel:    value[1].(uint8),
			}
			if err := db.FirstOrCreate(&event).Error; err != nil {
				return err
			}
		}
	}

	// Seed all report tags
	for key, value := range types.ReportTags {
		tag := &types.ReportTag{
			ID:          key,
			Description: value[0].(string),
			IsUser:      value[1].(bool),
			IsDeveloper: value[2].(bool),
			IsGame:      value[3].(bool),
		}
		if err := db.FirstOrCreate(&tag).Error; err != nil {
			return err
		}
	}

	// Seed the database with the test Developer and test Game IDs
	if err := db.FirstOrCreate(&types.Developer{
		Name: "Test Developer",
		ID:   "01HNPHQM5SPAG43J68R3NRX4M6",
	}).Error; err != nil {
		return err
	}

	var demogame *types.DeveloperGame = &types.DeveloperGame{
		Name:        "Test Game",
		ID:          "01HNPHRWS0N0AYMM5K4HN31V4W",
		DeveloperID: "01HNPHQM5SPAG43J68R3NRX4M6",
		Description: "This is a sample game provided by the server for testing use.",
		State:       1,
	}
	if err := db.FirstOrCreate(&demogame).Error; err != nil {
		return err
	}

	var tags []*types.FeatureTag
	for _, tag := range []string{
		"achievements",
		"controllers",
		"everyone",
		"legacy",
		"multidev",
		"oss",
		"onscratch",
		"ontw",
		"review",
		"save",
		"vchat",
	} {
		tags = append(tags, &types.FeatureTag{ID: tag})
	}
	db.Model(&demogame).Association("Features").Replace(tags)

	return nil
}

func ConvertDatabase(old_db *gorm.DB, new_db *gorm.DB, accounts_db *database.Database) error {

	// Step 1: Copy developers and games
	var old_developers []*old_types.Developers
	if err := old_db.Find(&old_developers).Error; err != nil {
		return err
	}

	log.Info("[1/3] Converting", len(old_developers), "developers and games...")
	for i, developer := range old_developers {

		// Copy developer
		new_developer := &types.Developer{
			ID:          developer.ID,
			Name:        developer.Name,
			CreatedAt:   time.Unix(developer.Created, 0),
			Description: developer.Description,
			State:       developer.State,
		}
		if err := new_db.FirstOrCreate(&new_developer).Error; err != nil {
			return err
		}

		// Copy games
		var old_games []*old_types.Games
		if err := old_db.Where("developerid = ?", developer.ID).Find(&old_games).Error; err != nil {
			return err
		}

		log.Debug("[", i+1, "/", len(old_developers), "] Found ", len(old_games), " games for developer ", developer.Name, ".")
		for _, game := range old_games {
			if err := new_db.FirstOrCreate(&types.DeveloperGame{
				ID:          game.ID,
				Name:        game.Name,
				DeveloperID: developer.ID,
				Description: "",
				CreatedAt:   game.Created,
				State:       game.State,
			}).Error; err != nil {
				return err
			}
		}
	}
	log.Info("Done converting developers and games.")
	time.Sleep(3 * time.Second)

	// Step 2: Convert users and saves
	var old_users []*old_types.Users
	if err := old_db.Find(&old_users).Error; err != nil {
		return err
	}

	log.Info("[2/3] Converting", len(old_users), "users and their saves...")
	for i, user := range old_users {

		// Generate a secret
		secret, err := accounts_db.CreateUserSecret()
		if err != nil {
			return err
		}

		// Create the user
		new_user := &types.User{
			ID:        user.ID,
			Username:  user.Username,
			Secret:    secret,
			Email:     user.Email,
			Password:  user.Password,
			State:     user.State,
			CreatedAt: time.Unix(user.Created, 0),
		}
		if err := new_db.FirstOrCreate(&new_user).Error; err != nil {
			return err
		}

		// Copy saves
		var old_saves []*old_types.Saves
		if err := old_db.Where("userid = ?", user.ID).Find(&old_saves).Error; err != nil {
			return err
		}

		log.Info("[", i+1, "/", len(old_users), "] Found ", len(old_saves), " saves for user ", user.Username, "...")
		for _, save := range old_saves {

			log.Debug(" > ", save.GameID, " (", save.SlotID, ")...")

			// Encrypt the save
			encrypted_save, err := accounts_db.Encrypt(new_user, save.Contents)
			if err != nil {
				return err
			}

			// Store the save
			new_save := &types.UserGameSave{
				UserID:          save.UserID,
				DeveloperGameID: save.GameID,
				SaveSlot:        save.SlotID,
				SaveData:        encrypted_save,
			}
			if err := new_db.Save(&new_save).Error; err != nil {
				return err
			}
		}
	}
	log.Info("Done converting users and saves.")
	time.Sleep(3 * time.Second)

	// Step 3: Convert memberships

	var old_developer_members []*old_types.DeveloperMembers
	if err := old_db.Find(&old_developer_members).Error; err != nil {
		return err
	}

	log.Info("[3/3] Converting", len(old_developer_members), "developer memberships...")
	for i, membership := range old_developer_members {
		log.Info("[", i+1, "/", len(old_developer_members), "] User: ", membership.UserID, " -> Developer: ", membership.DeveloperID, "...")
		new_db.Model(&types.Developer{ID: membership.DeveloperID}).Association("DeveloperMembers").Append(&types.DeveloperMember{
			UserID:      membership.UserID,
			DeveloperID: membership.DeveloperID,
		})
	}

	log.Info("Conversion complete.")
	time.Sleep(3 * time.Second)
	return nil
}
