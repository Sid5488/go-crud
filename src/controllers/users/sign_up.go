package controller

import (
	"fmt"

	rest_err "github.com/Sid5488/go-crud/src/configurations"
	"github.com/Sid5488/go-crud/src/models/request"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()),
		)

		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}
