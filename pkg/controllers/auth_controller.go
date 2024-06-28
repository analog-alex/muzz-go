package controllers

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/entities"
	"muzz-service/pkg/entities/cryptography"
	"muzz-service/pkg/repository"
	"net/http"
)

func Login(c *gin.Context) {

	// get login payload from request body
	var credentials entities.UserCredentials
	if err := c.Bind(&credentials); err != nil {
		entities.ErrResp(c, http.StatusBadRequest, "invalid payload", nil)
		return
	}

	// get user by username
	user, isPresent := repository.GetByEmail(credentials.Email)
	if !isPresent {
		entities.ErrResp(c, http.StatusNotFound, "user not found", nil)
		return
	}

	// validate password with bcrypt
	if !cryptography.CheckPasswordHash(credentials.Password, user.Password) {
		entities.ErrResp(c, http.StatusUnauthorized, "invalid credentials", nil)
		return
	}

	token, err := cryptography.GenerateJWToken(string(rune(user.ID)))

	if err != nil {
		entities.ErrResp(
			c,
			http.StatusInternalServerError,
			"invalid credentials",
			&entities.Extras{"id": user.ID},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
