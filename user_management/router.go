package user_management

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserData struct {
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
}

var users = map[string]UserData{
	"1": {
		Name:       "John",
		FamilyName: "Doe",
	},
	"2": {
		Name:       "Doe",
		FamilyName: "John",
	},
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
	user := users[id]
	if user == (UserData{}) {
		ctx.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})

}

// @BasePath /api/v1
// @Summary Create user
// @Description Create user
// @Tags Users
// @Accept json
// @Produce json
// @Param userData body UserData true "User data"
// @SUCCESS 200 {object} string "User created successfully"
// @Router /api/v1/users [post]
func createUser(ctx *gin.Context) {
	var user UserData
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	users[strconv.Itoa(len(users)+1)] = user
	fmt.Printf("Users: %v\n", users)
	ctx.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

func CreateUsersRouter(router *gin.RouterGroup) {
	router.GET("/users", getUsers)
	router.POST("/user/:id", getUser)
	router.POST("/users", createUser)
}
