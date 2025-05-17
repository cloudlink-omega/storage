package common

import (
	"github.com/cloudlink-omega/storage/pkg/types"
	"github.com/oklog/ulid/v2"

	"gorm.io/gorm"
)

func LogEvent(db *gorm.DB, event any) string {

	switch my_event := event.(type) {

	case *types.UserEvent:
		my_event.ID = ulid.Make().String()
		if err := db.Create(my_event).Error; err != nil {
			panic(err)
		}
		return my_event.ID

	case *types.DeveloperEvent:
		if err := db.Create(my_event).Error; err != nil {
			panic(err)
		}
		return my_event.ID

	case *types.SystemEvent:
		if err := db.Create(my_event).Error; err != nil {
			panic(err)
		}
		return my_event.ID

	}

	panic("invalid event type")
}
