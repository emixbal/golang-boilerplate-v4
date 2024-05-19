package customer_test

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"golang-websocket/api/mocks/customer"
// 	"golang-websocket/api/models"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestList(t *testing.T) {
// 	mockCustomerRepo := new(customer.CustomerRepositoryMock)

// 	mockCustomer := &models.Customer{
// 		Nim:   "23142008",
// 		Nama:  "Dadang",
// 		Kelas: "TIB",
// 	}

// 	mockListCustomer := make([]*models.Customer, 0)
// 	mockListCustomer = append(mockListCustomer, mockCustomer)

// 	t.Run("success", func(t *testing.T) {
// 		mockCustomerRepo.On("List", mock.Anything).Return(mockListCustomer, nil).Once()
// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		list, err := u.List(context.TODO())

// 		assert.NoError(t, err)
// 		assert.Len(t, list, len(mockListCustomer))
// 		mockCustomerRepo.AssertExpectations(t)
// 	})

// 	t.Run("error-failed", func(t *testing.T) {
// 		mockCustomerRepo.On("List", mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()
// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		list, err := u.List(context.TODO())

// 		assert.Error(t, err)
// 		assert.Len(t, list, 0)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})
// }

// func TestDetail(t *testing.T) {
// 	mockCustomerRepo := new(customer.CustomerRepositoryMock)

// 	mockCustomer := models.Customer{
// 		ID:    3,
// 		Nim:   "23142008",
// 		Nama:  "Dadang",
// 		Kelas: "TIB",
// 	}
// 	fmt.Println(mockCustomer)

// 	t.Run("success", func(t *testing.T) {
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockCustomer, nil).Once()

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*5)
// 		customer, err := u.Detail(context.TODO(), mockCustomer.ID)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, customer)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})

// 	t.Run("error-failed", func(t *testing.T) {
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)
// 		customer, err := u.Detail(context.TODO(), mockCustomer.ID)

// 		assert.Error(t, err)
// 		assert.Nil(t, customer)
// 		mockCustomerRepo.AssertExpectations(t)

// 	})
// }

// func TestInsert(t *testing.T) {
// 	mockCustomerRepo := new(customer.CustomerRepositoryMock)
// 	mockCustomer := models.Customer{
// 		Nim:   "23142008",
// 		Nama:  "Dadang",
// 		Kelas: "TIB",
// 	}

// 	t.Run("success", func(t *testing.T) {
// 		tempMockCustomer := mockCustomer
// 		tempMockCustomer.ID = 0
// 		mockCustomerRepo.On("Insert", mock.Anything, mock.Anything).Return(&mockCustomer, nil).Once()
// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		customer, err := u.Insert(context.TODO(), tempMockCustomer)
// 		assert.NoError(t, err)
// 		assert.Equal(t, mockCustomer.Nama, tempMockCustomer.Nama)
// 		assert.NotNil(t, &customer)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})

// 	t.Run("error-failed", func(t *testing.T) {
// 		mockCustomerRepo.On("Insert", mock.Anything, mock.Anything).Return(nil, errors.New("unexpected error")).Once()

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		customer, err := u.Insert(context.TODO(), mockCustomer)

// 		assert.Error(t, err)
// 		assert.Nil(t, customer)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	mockCustomerRepo := new(customer.CustomerRepositoryMock)
// 	// var mockMhs *models.Customer
// 	mockCustomers := models.Customer{
// 		Nim:   "23142008",
// 		Nama:  "Dadang",
// 		Kelas: "TIB",
// 	}
// 	mockCustomer := make(map[string]interface{})
// 	mockCustomer["nama"] = "Dadang"
// 	id := int(12)

// 	t.Run("success", func(t *testing.T) {

// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockCustomers, nil).Once()
// 		mockCustomerRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Once().Return(nil)
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockCustomers, nil).Once()

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		customer, err := u.Update(context.TODO(), mockCustomer, id)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, customer)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})

// 	t.Run("error-in-db", func(t *testing.T) {
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error"))

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		customer, err := u.Update(context.TODO(), mockCustomer, id)

// 		assert.Error(t, err)
// 		assert.Nil(t, customer)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	mockCustomerRepo := new(customer.CustomerRepositoryMock)
// 	mockCustomers := models.Customer{
// 		Nim:   "23142008",
// 		Nama:  "Dadang",
// 		Kelas: "TIB",
// 	}
// 	id := 3

// 	t.Run("success", func(t *testing.T) {
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(&mockCustomers, nil).Once()
// 		mockCustomerRepo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		err := u.Delete(context.TODO(), id)

// 		assert.NoError(t, err)

// 		mockCustomerRepo.AssertExpectations(t)
// 	})

// 	t.Run("customer-not-found", func(t *testing.T) {
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, nil).Once()
// 		mockCustomerRepo.On("Delete", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

// 		err := u.Delete(context.TODO(), id)

// 		assert.Nil(t, err)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})

// 	t.Run("error-in-db", func(t *testing.T) {
// 		mockCustomerRepo.On("Detail", mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error")).Once()
// 		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)
// 		err := u.Delete(context.TODO(), id)

// 		assert.Error(t, err)
// 		mockCustomerRepo.AssertExpectations(t)
// 	})
// }
