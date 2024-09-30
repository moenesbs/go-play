package main

import (
	database "example/hello/database"
	_ "example/hello/docs"
	userRouter "example/hello/user_management"
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
	userRouter.CreateUsersRouter(api)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	database.PopulateDatabase()
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
