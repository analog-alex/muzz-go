package controllers

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/repository"
	"muzz-service/pkg/types"
	"muzz-service/pkg/types/cryptography"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	var credentials types.UserCredentials
	if err := c.Bind(&credentials); err != nil {
		types.ErrResp(c, http.StatusBadRequest, "invalid payload", nil)
		return
	}

	// fetch user by email
	users, err := repository.GetUsersByEmail(credentials.Email)
	if err != nil || len(users) == 0 || len(users) > 1 {
		types.ErrResp(c, http.StatusNotFound, "user not found", nil)
		return
	}

	user := users[0]

	// validate incoming password with user password
	if !cryptography.CheckPasswordHash(user.Password, credentials.Password) {
		types.ErrResp(c, http.StatusUnauthorized, "invalid credentials", nil)
		return
	}

	token, err := cryptography.GenerateJWToken(strconv.Itoa(user.ID))

	if err != nil {
		types.ErrResp(
			c,
			http.StatusInternalServerError,
			"token generation failed",
			&types.Extras{"id": user.ID, "reason": err.Error()},
		)
		return
	}

	types.OkResp(c, http.StatusOK, types.LoginResponse{Token: token})
}
