package userRouter

import (
	"example/hello/database"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
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
	users := database.GetUsers()
	context.JSON(200, gin.H{
		"users": users,
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
	fmt.Printf("Users: %v\n", user)
	userId := database.InsertUser(database.UserStruct{Name: user.Name, FamilyName: user.FamilyName})
	ctx.JSON(200, gin.H{
		"message": "User created successfully, id: " + fmt.Sprint(userId),
	})
}

func CreateUsersRouter(router *gin.RouterGroup) {
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
}
