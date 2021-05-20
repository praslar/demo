package model

import (
	postgres "development-kit/pkg/db/postgres"
	"fmt"
	log "github.com/go-kit/kit/log"
	uuid "github.com/google/uuid"
	"time"
)

var logger log.Logger

// DevelopmentKit describes the structure.
type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatorID uuid.UUID  `json:"creator_id"`
	UpdaterID uuid.UUID  `json:"updater_id"`
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

// HealthCheckResponse a json encoded health check response
type HealthCheckResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

func AutoMigration() (err error) {
	dbPublic, err := postgres.GetDatabase("default")
	if err != nil {
		logger.Log("err connect DB: ", err)
		return err
	}
	_, err = dbPublic.DB().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err != nil {
		return fmt.Errorf("error while creating DB extension 'uuid-ossp': %s", err)
	}
	t := dbPublic.AutoMigrate()
	return t.Error
}
