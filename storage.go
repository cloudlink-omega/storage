package storage

import "gorm.io/gorm"

/*
 * Upgrades the old database to the new one implemented using GORM. Since migrations aren't capable of
 * upgrading the database, this function will need to be called to upgrade the database.
 * TODO.
 */
func UpgradeToV1(oldDB *gorm.DB, newDB *gorm.DB) error {
	// TODO
	return nil
}
