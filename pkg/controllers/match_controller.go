package controllers

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/repository"
	"muzz-service/pkg/types"
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
	users, err := repository.GetAllUsersExcluding(userId)

	if err != nil {
		types.ErrResp(c, 500, "error fetching users", nil)
		return
	}

	types.OkResp(c, 200, users)
}

func Swipe(c *gin.Context) {
	// TODO
	// 1. create swipe
	// 2. check if other direction swipe was a right swipe
	// 3. if both swipes are right, create a match
}
