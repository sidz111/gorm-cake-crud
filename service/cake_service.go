package service

import (
	"context"
	"fmt"

	"github.com/sidz111/cake-gorm/model"
	"github.com/sidz111/cake-gorm/repository"
)

type CakeService interface {
	Create(ctx context.Context, cake *model.Cake) error
	GetById(ctx context.Context, id int) (*model.Cake, error)
	GetAll(ctx context.Context) ([]model.Cake, error)
	Update(ctx context.Context, cake *model.Cake) error
	Delete(ctx context.Context, id int) error
}

type cakeService struct {
	repo repository.CakeRepository
}

func NewCakeService(repo repository.CakeRepository) CakeService {
	return &cakeService{repo: repo}
}

func (s *cakeService) Create(ctx context.Context, cake *model.Cake) error {
	if err := CakeValidation(cake); err != nil {
		return err
	}
	return s.repo.Create(ctx, cake)
}
func (s *cakeService) GetById(ctx context.Context, id int) (*model.Cake, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id %d", id)
	}
	return s.repo.GetById(ctx, id)
}
func (s *cakeService) GetAll(ctx context.Context) ([]model.Cake, error) {
	return s.repo.GetAll(ctx)
}
func (s *cakeService) Update(ctx context.Context, cake *model.Cake) error {
	if err := CakeValidation(cake); err != nil {
		return err
	}
	return s.repo.Update(ctx, cake)
}
func (s *cakeService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid id %d", id)
	}
	return s.repo.Delete(ctx, id)
}

func CakeValidation(cake *model.Cake) error {
	if cake == nil {
		return fmt.Errorf("cake cannot be nil")
	}

	if cake.Name == "" {
		return fmt.Errorf("please enter name")
	}

	if cake.Price <= 0 {
		return fmt.Errorf("price should be positive")
	}

	return nil
}
