package service

import (
	"context"

	"github.com/sidz111/cake-gorm/model"
	"github.com/sidz111/cake-gorm/repository"
)

type CakeService interface {
	Create(ctx context.Context, cake *model.Cake) error
	GetById(ctx context.Context, id int) (*model.Cake, error)
	GetAll(ctx context.Context) (*[]model.Cake, error)
	Update(ctx context.Context, cake *model.Cake) error
	Delete(ctx context.Context, id int) error
}

type cakeService struct {
	repo repository.CakeRepository
}

// func NewCakeRepository(repo repository.CakeRepository) CakeRepository {
// 	return &cakeService(repo)
// }

// func (r *Cake) Create(ctx context.Context, cake *model.Cake) error
// 	GetById(ctx context.Context, id int) (*model.Cake, error)
// 	GetAll(ctx context.Context) (*[]model.Cake, error)
// 	Update(ctx context.Context, cake *model.Cake) error
// 	Delete(ctx context.Context, id int) error
