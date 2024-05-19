package customer

import (
	"context"
	"golang-websocket/api/database"
	"golang-websocket/api/helper"
	"golang-websocket/api/models"
	"golang-websocket/api/repository/customer"
	"golang-websocket/api/usecase"
	ucase "golang-websocket/api/usecase/customer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type CustomerHandler struct {
	CustomerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler() CustomerHandler {
	timeout := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	db := database.Load()
	repoCustomer := customer.NewCustomerRepository(db)
	ucaseCustomer := ucase.NewCustomerUsecase(repoCustomer, timeout)
	return CustomerHandler{
		CustomerUsecase: ucaseCustomer,
	}
}

func (m *CustomerHandler) List(c *gin.Context) {
	var res = c.Writer
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	customer, err := m.CustomerUsecase.List(ctx)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Responses(res, http.StatusOK, "Success", customer)
}

func (m *CustomerHandler) Detail(c *gin.Context) {
	var res = c.Writer
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	customer, err := m.CustomerUsecase.Detail(ctx, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}
	helper.Responses(res, http.StatusOK, "Success", customer)
}

func (m *CustomerHandler) Insert(c *gin.Context) {
	var customer models.Customer
	var res = c.Writer
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	// var validators *validator.Validate
	// config := &validator.Config{TagName: "validate"}
	// validators = validator.New(config)
	// err := validators.Struct(customer)

	// if err != nil {
	// 	helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
	// 	return
	// }

	result, err := m.CustomerUsecase.Insert(ctx, customer)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}
	helper.Responses(res, http.StatusOK, "Success", result)
}

func (m *CustomerHandler) Update(c *gin.Context) {
	var datas = make(map[string]interface{})
	var res = c.Writer
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	datas["nama"] = c.Request.FormValue("nama")
	datas["nim"] = c.Request.FormValue("nim")
	datas["kelas"] = c.Request.FormValue("kelas")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	customer, err := m.CustomerUsecase.Update(ctx, datas, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}

	helper.Responses(res, http.StatusOK, "Succes", customer)
}

func (m *CustomerHandler) Delete(c *gin.Context) {
	var res = c.Writer
	id, err := helper.ToInt(c.Param("id"))
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = m.CustomerUsecase.Delete(ctx, id)
	if err != nil {
		helper.HandlerErrorQuery(res, err)
		return
	}

	helper.Responses(res, http.StatusOK, "Success", "Data Telah Dihapus")
}
