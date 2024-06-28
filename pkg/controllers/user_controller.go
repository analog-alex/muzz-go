package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/repository"
	"muzz-service/pkg/types"
	"muzz-service/pkg/types/cryptography"
	"net/http"
)

func Create(c *gin.Context) {
	user := types.User{
		Email:    types.GenerateEmail(),
		Password: types.GeneratePassword(),
		Name:     types.GenerateName(),
		Gender:   types.GenerateGender(),
		Age:      types.GenerateAge(),
	}

	// hash the password but keep the value of original password
	password := user.Password
	hashedPassword, err := cryptography.HashPassword(user.Password)
	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error hashing password", nil)
		return
	}

	user.Password = hashedPassword
	persistedUser, err := repository.CreateUser(user)
	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error creating user", nil)
		return
	}

	// important: return the original password
	persistedUser.Password = password
	types.OkResp(c, http.StatusCreated, persistedUser)
}

func GetAll(c *gin.Context) {
	users, err := repository.GetAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		types.ErrResp(c, http.StatusInternalServerError, "error fetching users", nil)
		return
	}

	types.OkResp(c, http.StatusOK, users)
}
