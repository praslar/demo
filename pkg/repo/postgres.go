package repo

import (
	"context"
	"github.com/jinzhu/gorm"
)

func NewBasicPostgresDatabase(db *gorm.DB) PostgresDatabase {
	return &basicPostgresDatabase{
		db: db,
	}
}

type basicPostgresDatabase struct {
	db *gorm.DB
}

func (b basicPostgresDatabase) Create(ctx context.Context, req interface{}) error {
	panic("implement me")
}

func (b basicPostgresDatabase) Update(ctx context.Context, req interface{}) error {
	return b.db.Save(req).Error
}

func (b basicPostgresDatabase) GetOneByID(ctx context.Context, id string) (r interface{}, e error) {
	if err := b.db.Where("id = ?", id).First(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (b basicPostgresDatabase) Delete(ctx context.Context, req interface{}) error {
	return b.db.Delete(req).Error
}

func (b basicPostgresDatabase) Get() *gorm.DB {
	return b.db
}

type PostgresDatabase interface {
// Add your db method here,
	Create(ctx context.Context, req interface{}) error
	Update(ctx context.Context, req interface{}) error
	GetOneByID(ctx context.Context, id string) (r interface{}, e error)
	Delete(ctx context.Context, req interface{}) error
	Get() *gorm.DB
}