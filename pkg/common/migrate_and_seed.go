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
		&types.Developer{},
		&types.DeveloperEvent{},
		&types.DeveloperGame{},
		&types.DeveloperMember{},
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
		Features: []*types.FeatureTag{
			{ID: "achievements"},
			{ID: "controllers"},
			{ID: "everyone"},
			{ID: "legacy"},
			{ID: "multidev"},
			{ID: "oss"},
			{ID: "onscratch"},
			{ID: "ontw"},
			{ID: "review"},
			{ID: "save"},
			{ID: "vchat"},
		},
	}).Error; err != nil {
		return err
	}

	return nil
}

func ConvertDatabase(old_db *gorm.DB, new_db *gorm.DB, accounts_db *database.Database) error {

	// Step 1: Convert users
	var old_users []*old_types.Users
	if err := old_db.Find(&old_users).Error; err != nil {
		return err
	}

	log.Debug("[1/5] Converting", len(old_users), "users...")
	for _, user := range old_users {

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
	}

	// Step 2: Copy developers
	var old_developers []*old_types.Developers
	if err := old_db.Find(&old_developers).Error; err != nil {
		return err
	}

	log.Debug("[2/5] Converting", len(old_developers), "developers...")
	for _, developer := range old_developers {
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
	}

	// Step 3: Copy games
	var old_games []*old_types.Games
	if err := old_db.Find(&old_games).Error; err != nil {
		return err
	}

	log.Debug("[3/5] Converting", len(old_games), "games...")
	for _, game := range old_games {
		new_game := &types.DeveloperGame{
			ID:          game.ID,
			Name:        game.Name,
			DeveloperID: game.DeveloperID,
			Description: "",
			CreatedAt:   game.Created,
			State:       game.State,
		}
		if err := new_db.FirstOrCreate(&new_game).Error; err != nil {
			return err
		}
	}

	// Step 4: Copy developer memberships
	var old_developer_members []*old_types.DeveloperMembers
	if err := old_db.Find(&old_developer_members).Error; err != nil {
		return err
	}

	log.Debug("[4/5] Converting", len(old_developer_members), "developer memberships...")
	for _, old_member := range old_developer_members {
		new_member := &types.DeveloperMember{
			DeveloperID: old_member.DeveloperID,
			UserID:      old_member.UserID,
		}
		if err := new_db.FirstOrCreate(&new_member).Error; err != nil {
			return err
		}
	}

	// Step 5: Copy saves
	var old_saves []*old_types.Saves
	if err := old_db.Find(&old_saves).Error; err != nil {
		return err
	}

	log.Debug("[5/5] Converting", len(old_saves), "saves. This may take a while...")
	for _, save := range old_saves {

		// We will need to get the user's secret to encrypt the save
		user, err := accounts_db.GetUser(save.UserID)
		if err != nil {
			return err
		}

		// Encrypt the save
		encrypted_save, err := accounts_db.Encrypt(user, save.Contents)
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

	log.Debug("Conversion complete.")
	return nil
}
