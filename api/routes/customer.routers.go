package routes

import (
	"golang-websocket/api/controllers/customer"

	"github.com/gin-gonic/gin"
)

func RouteCustomer(route *gin.RouterGroup) {

	handlerCustomer := customer.NewCustomerHandler()
	router := route.Group("/customer")
	{
		// router.Use(middleware.MiddlewareAuthentication)
		{
			router.GET("/list", handlerCustomer.List)
			router.GET("/detail/:id", handlerCustomer.Detail)
			router.POST("/insert", handlerCustomer.Insert)
			router.PUT("/update/:id", handlerCustomer.Update)
			router.DELETE("/delete/:id", handlerCustomer.Delete)
		}
	}
}
