package handler

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/dao"
	"muzz-service/pkg/types"
	"net/http"
	"strconv"
)

func Discover(c *gin.Context) {
	// as this is a protected route, we can get the userId from the context
	userCtx, _ := c.Get("userId")
	userId, err := strconv.Atoi(userCtx.(string))
	if err != nil {
		types.ErrResp(c, 500, "error fetching user id from context", nil)
		return
	}

	// get all users except the current user
	users, err := dao.GetAllUsersExcludingSwipes(userId)

	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error fetching users", nil)
		return
	}

	// convert to user profiles
	profiles := make([]types.UserProfile, len(users))
	for i, user := range users {
		profiles[i] = user.ToUserProfile()
	}

	types.OkResp(c, http.StatusOK, profiles)
}
