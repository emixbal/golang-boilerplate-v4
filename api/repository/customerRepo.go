package repository

import (
	"context"
	"golang-websocket/api/models"
)

type CustomerRepository interface {
	List(ctx context.Context) ([]*models.Customer, error)
	Detail(ctx context.Context, id int) (*models.Customer, error)
	Insert(ctx context.Context, Customer models.Customer) (*models.Customer, error)
	Update(ctx context.Context, datas map[string]interface{}, id int) error
	Delete(ctx context.Context, id int) error
}
