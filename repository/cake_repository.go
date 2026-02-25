package repository

import (
	"context"

	"github.com/sidz111/cake-gorm/model"
	"gorm.io/gorm"
)

type CakeRepository interface {
	Create(ctx context.Context, cake *model.Cake) error
	GetById(ctx context.Context, id int) (*model.Cake, error)
	GetAll(ctx context.Context) ([]model.Cake, error)
	Update(ctx context.Context, cake *model.Cake) error
	Delete(ctx context.Context, id int) error
}

type cakeRepo struct {
	DB *gorm.DB
}

func NewCakeRepository(db *gorm.DB) CakeRepository {
	return &cakeRepo{DB: db}
}

func (r *cakeRepo) Create(ctx context.Context, cake *model.Cake) error {
	err := r.DB.WithContext(ctx).Create(cake).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *cakeRepo) GetById(ctx context.Context, id int) (*model.Cake, error) {
	var cake *model.Cake
	err := r.DB.WithContext(ctx).First(&cake, id).Error
	if err != nil {
		return nil, err
	}
	return cake, nil
}
func (r *cakeRepo) GetAll(ctx context.Context) ([]model.Cake, error) {
	var cakes []model.Cake
	err := r.DB.WithContext(ctx).Find(&cakes).Error
	if err != nil {
		return nil, err
	}
	return cakes, nil
}
func (r *cakeRepo) Update(ctx context.Context, cake *model.Cake) error {
	err := r.DB.WithContext(ctx).Save(cake).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *cakeRepo) Delete(ctx context.Context, id int) error {
	err := r.DB.WithContext(ctx).Delete(&model.Cake{}, id).Error
	if err != nil {
		return err
	}
	return nil

}
