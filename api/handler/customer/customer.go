package customer

import (
	"golang-websocket/api/database"
	"golang-websocket/api/helper"
	"golang-websocket/api/models"
	"golang-websocket/api/repository/customer"
	"golang-websocket/api/usecase"
	ucase "golang-websocket/api/usecase/customer"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
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

func (h *CustomerHandler) List(c *fiber.Ctx) error {
	ctx := c.Context()
	customer, err := h.CustomerUsecase.List(ctx)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    customer,
	})
}

func (h *CustomerHandler) Detail(c *fiber.Ctx) error {
	id, err := helper.ToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := c.Context()
	customer, err := h.CustomerUsecase.Detail(ctx, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    customer,
	})
}

func (h *CustomerHandler) Insert(c *fiber.Ctx) error {
	var customer models.Customer
	ctx := c.Context()

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result, err := h.CustomerUsecase.Insert(ctx, customer)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    result,
	})
}

func (h *CustomerHandler) Update(c *fiber.Ctx) error {
	id, err := helper.ToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var datas = make(map[string]interface{})
	if err := c.BodyParser(&datas); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := c.Context()
	customer, err := h.CustomerUsecase.Update(ctx, datas, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    customer,
	})
}

func (h *CustomerHandler) Delete(c *fiber.Ctx) error {
	id, err := helper.ToInt(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx := c.Context()
	err = h.CustomerUsecase.Delete(ctx, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    "Data Telah Dihapus",
	})
}
