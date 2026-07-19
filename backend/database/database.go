package database

import (
	"log"

	"pelacakan-fruit-transport/config"
	"pelacakan-fruit-transport/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens a PostgreSQL connection via GORM, runs AutoMigrate,
// and returns the database handle.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the Transport model to create / update the transports table.
	if err := db.AutoMigrate(&model.Transport{}); err != nil {
		return nil, err
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}
