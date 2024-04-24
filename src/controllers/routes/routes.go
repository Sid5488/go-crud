package routes

import (
	controller "github.com/Sid5488/go-crud/src/controllers/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/users/get-user-by-id/:id", controller.FindUserById)
	r.GET("/users/get-user-by-email/:email", controller.FindUserByEmail)
	r.POST("/users/sign-up", controller.SignUp)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
