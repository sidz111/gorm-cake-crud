package service

import (
	"context"

	"github.com/sidz111/cake-gorm/model"
)

type CakeService interface {
	Create(ctx context.Context, cake *model.Cake) error
}
