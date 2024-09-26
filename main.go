package main

import (
	_ "example/hello/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func healthCheck(g *gin.Context) {
	g.JSON(http.StatusOK, "all good")
}

func main() {
	router := gin.Default()
	api := router.Group("/api/v1")
	router.GET("/", healthCheck)
	HelloRouter(api)
	UsersRouter(api)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
