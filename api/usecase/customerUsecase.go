package usecase

import (
	"context"
	"golang-websocket/api/models"
)

type CustomerUsecase interface {
	List(ctx context.Context) ([]*models.Customer, error)
	Detail(ctx context.Context, id int) (*models.Customer, error)
	Insert(ctx context.Context, mahasiswa models.Customer) (*models.Customer, error)
	Update(ctx context.Context, datas map[string]interface{}, id int) (*models.Customer, error)
	Delete(ctx context.Context, id int) error
}
