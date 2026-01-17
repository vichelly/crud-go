package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vichelly/crud-go/src/config/validation"
	"github.com/vichelly/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr.Message)
		return
	}
	fmt.Println(userRequest)
}
