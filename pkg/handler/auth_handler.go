package handler

import (
	"github.com/gin-gonic/gin"
	cryptography2 "muzz-service/pkg/cryptography"
	"muzz-service/pkg/dao"
	"muzz-service/pkg/types"
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
	users, err := dao.GetUsersByEmail(credentials.Email)
	if err != nil || len(users) == 0 || len(users) > 1 {
		types.ErrResp(c, http.StatusNotFound, "user not found", nil)
		return
	}

	user := users[0]

	// validate incoming password with user password
	if !cryptography2.CheckPasswordHash(user.Password, credentials.Password) {
		types.ErrResp(c, http.StatusUnauthorized, "invalid credentials", nil)
		return
	}

	token, err := cryptography2.GenerateJWToken(strconv.Itoa(user.ID))

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
