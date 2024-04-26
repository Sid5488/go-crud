package routes

import (
	controllers "github.com/Sid5488/go-crud/src/controllers/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controllers.UserControllerInterface,
) {
	r.GET("/users/get-user-by-id/:id", userController.FindUserById)
	r.GET("/users/get-user-by-email/:email", userController.FindUserByEmail)
	r.POST("/users/sign-up", userController.SignUp)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
