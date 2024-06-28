package controllers

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/entities"
	"net/http"
)

var users []entities.User

func Create(c *gin.Context) {

	user := entities.User{
		ID:       len(users) + 1,
		Email:    entities.GenerateEmail(),
		Password: entities.GeneratePassword(),
		Name:     entities.GenerateName(),
		Gender:   entities.GenerateGender(),
		Age:      entities.GenerateAge(),
	}

	users = append(users, user)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusCreated, user)
}

func GetAll(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, users)
}
