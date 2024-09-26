package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	UserId string `json:"id"`
	Name   string `json:"name"`
}

var users = []UserData{
	{
		UserId: "1",
		Name:   "John",
	},
	{
		UserId: "2",
		Name:   "Doe",
	},
}

func FilterUsersById(users []UserData, userId string) UserData {
	var userInfo UserData
	for _, user := range users {
		if user.UserId == userId {
			userInfo = user
		}
	}
	return userInfo
}

// @BasePath /api/v1
// @Summary Get all users
// @Description Get all users
// @Tags Users
// @Accept json
// @Produce json
// @SUCCESS 200
// @Router /api/v1/users [get]
func getUsers(context *gin.Context) {
	fmt.Printf("Users: %v\n", users)
	context.IndentedJSON(200, users)
}

// @BasePath /api/v1
// @Summary Get user by id
// @Description Get user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @SUCCESS 200 {object} string "User id"
// @Router /api/v1/user/{id} [post]
func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"message": "User id is required",
		})
		return
	}
	fmt.Printf("User id: %s\n", id)
	ctx.JSON(200, gin.H{
		"user": FilterUsersById(users, id),
	})

}

func UsersRouter(router *gin.RouterGroup) {
	router.GET("/users", getUsers)
	router.POST("/user/:id", getUser)
}
