package routes

import (
	"golang-websocket/api/handler/mahasiswa"
	"golang-websocket/api/middleware"

	"github.com/gin-gonic/gin"
)

func RouteMahasiswa(route *gin.RouterGroup) {

	handlerMahasiswa := mahasiswa.NewMahasiswaHandler()
	router := route.Group("/mahasiswa")
	{
		router.Use(middleware.MiddlewareAuthentication)
		{
			router.GET("/list", handlerMahasiswa.List)
			router.GET("/detail/:id", handlerMahasiswa.Detail)
			router.POST("/insert", handlerMahasiswa.Insert)
			router.PUT("/update/:id", handlerMahasiswa.Update)
			router.DELETE("/delete/:id", handlerMahasiswa.Delete)
		}
	}
}
