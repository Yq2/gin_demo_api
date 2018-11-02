package router

import (
	. "gin_demo_api/apis"
	"gopkg.in/gin-gonic/gin.v1"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/person/:id", DelPersonApi)

	return router
}
