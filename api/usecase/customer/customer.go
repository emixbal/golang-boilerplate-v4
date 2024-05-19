package customer

import (
	"context"
	"golang-websocket/api/models"
	"golang-websocket/api/repository"
	"golang-websocket/api/usecase"
	"time"
)

type customerUsecase struct {
	customerRepo   repository.CustomerRepository
	contextTimeout time.Duration
}

func NewCustomerUsecase(customerRepo repository.CustomerRepository, timeout time.Duration) usecase.CustomerUsecase {
	return &customerUsecase{
		customerRepo:   customerRepo,
		contextTimeout: timeout,
	}
}

func (u *customerUsecase) List(c context.Context) ([]*models.Customer, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	customer, err := u.customerRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (u *customerUsecase) Detail(c context.Context, id int) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	customer, err := u.customerRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (u *customerUsecase) Insert(c context.Context, customer models.Customer) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	customerResult, err := u.customerRepo.Insert(ctx, customer)
	if err != nil {
		return nil, err
	}
	return customerResult, nil
}

func (u *customerUsecase) Update(c context.Context, datas map[string]interface{}, id int) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.customerRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	err = u.customerRepo.Update(ctx, datas, id)
	if err != nil {
		return nil, err
	}

	customers, err := u.customerRepo.Detail(ctx, id)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (u *customerUsecase) Delete(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.customerRepo.Detail(ctx, id)
	if err != nil {
		return err
	}

	err = u.customerRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
