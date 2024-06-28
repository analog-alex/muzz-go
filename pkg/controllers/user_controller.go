package controllers

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/entities"
	"muzz-service/pkg/entities/cryptography"
	"muzz-service/pkg/repository"
	"net/http"
)

func Create(c *gin.Context) {
	user := entities.User{
		Email:    entities.GenerateEmail(),
		Password: entities.GeneratePassword(),
		Name:     entities.GenerateName(),
		Gender:   entities.GenerateGender(),
		Age:      entities.GenerateAge(),
	}

	// hash the password but keep the value of original password
	password := user.Password
	hashedPassword, err := cryptography.HashPassword(user.Password)
	if err != nil {
		entities.ErrResp(c, http.StatusInternalServerError, "error hashing password", nil)
		return
	}

	user.Password = hashedPassword
	persistedUser := repository.Create(user)

	// important: return the original password
	persistedUser.Password = password
	entities.OkResp(c, http.StatusCreated, persistedUser)
}

func GetAll(c *gin.Context) {
	users := repository.GetAll()
	entities.OkResp(c, http.StatusOK, users)
}
