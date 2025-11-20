package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jpriscio/first-CRUD/src/configuration/logger"
	"github.com/jpriscio/first-CRUD/src/configuration/validation"
	"github.com/jpriscio/first-CRUD/src/controller/model/request"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	logger.Info("Init Create User",
		zap.String("journey", "createUser"))

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("Erro tryng to validate user info", err,
			zap.String("journey", "createUser"))
		rest_Err := validation.ValidateUserError(err)

		c.JSON(rest_Err.Code, rest_Err)
		return

	}

	c.JSON(200, gin.H{
		"message": "Usuário criado com sucesso!",
		"email":   UserRequest.Email,
		"name":    UserRequest.Name,
		"age":     UserRequest.Age,
	})

	fmt.Println(UserRequest)
}
