package service

import (
	"context"
	"development-kit/pkg/model"
	"development-kit/pkg/repo"
)

// DevelopmentKitService describes the service.
type DevelopmentKitService interface {
	// Add your methods here, e.x:
	Status(ctx context.Context, name, version string) (rs model.HealthCheckResponse, err error)
	CreateDemo(ctx context.Context, req model.DemoCreateRequest) (rs string, err error)
	UpdateDemo(ctx context.Context, req model.DemoUpdateRequest) (rs string, err error)
	GetOneDemo(ctx context.Context, req model.DemoGetRequest) (rs string, err error)
	DeleteDemo(ctx context.Context, req model.DemoDeleteRequest) (rs string, err error)
}

type basicDevelopmentKitService struct {
	pgDB repo.PostgresDatabase
}

func (b *basicDevelopmentKitService) CreateDemo(ctx context.Context, req model.DemoCreateRequest) (rs string, err error) {
	// TODO implement the business logic of CreateDemo
	return rs, err
}
func (b *basicDevelopmentKitService) UpdateDemo(ctx context.Context, req model.DemoUpdateRequest) (rs string, err error) {
	// TODO implement the business logic of UpdateDemo
	return rs, err
}
func (b *basicDevelopmentKitService) GetOneDemo(ctx context.Context, req model.DemoGetRequest) (rs string, err error) {
	// TODO implement the business logic of GetOneDemo
	return rs, err
}
func (b *basicDevelopmentKitService) DeleteDemo(ctx context.Context, req model.DemoDeleteRequest) (rs string, err error) {
	// TODO implement the business logic of DeleteDemo
	return rs, err
}

// NewBasicDevelopmentKitService returns a naive, stateless implementation of DevelopmentKitService.
func NewBasicDevelopmentKitService(db repo.PostgresDatabase) DevelopmentKitService {
	return &basicDevelopmentKitService{pgDB: db}
}

// New returns a DevelopmentKitService with all of the expected middleware wired in.
func New(middleware []Middleware, db repo.PostgresDatabase) DevelopmentKitService {
	var svc DevelopmentKitService = NewBasicDevelopmentKitService(db)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicDevelopmentKitService) Status(ctx context.Context, name string, version string) (rs model.HealthCheckResponse, err error) {
	return rs, err
}
