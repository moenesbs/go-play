package main

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Welcome router
// @Description This is the welcome router for users
// @Tags Welcome
// @Accept json
// @Produce json
// @SUCCESS 200 {object} string "ok"
// @Router /api/v1/hello [get]
func Helloworld(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World from Example service",
	})
}

// @BasePath /api/v1
// @Summary Welcome router
// @Description This is the welcome router for users
// @Tags Welcome
// @Accept json
// @Produce json
// @SUCCESS 200 {object} string "ok"
// @Router /api/v1/goodbye [get]
func Goodbye(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Goodbye from Example service",
	})
}

func HelloRouter(router *gin.RouterGroup) {
	router.GET("/helloworld", Helloworld)
	router.GET("/goodbye", Goodbye)
}
