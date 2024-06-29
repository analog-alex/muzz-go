package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"muzz-service/pkg/dao"
	"muzz-service/pkg/types"
	"net/http"
	"strconv"
)

func Swipe(c *gin.Context) {
	var swipe types.SwipeRequest
	if err := c.Bind(&swipe); err != nil {
		types.ErrResp(c, http.StatusBadRequest, "invalid payload", nil)
		return
	}

	// as this is a protected route, we can get the userId from the context
	userCtx, _ := c.Get("userId")
	userId, err := strconv.Atoi(userCtx.(string))
	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error fetching user id from context", nil)
		return
	}

	// check if user is not swiping themselves
	if userId == swipe.UserId {
		types.ErrResp(c, http.StatusBadRequest, "cannot swipe yourself", nil)
		return
	}

	// create swipe
	swipeModel := types.Swipe{
		Accept: swipe.Preference == "YES",
		From:   userId,
		To:     swipe.UserId,
	}

	_, err = dao.CreateSwipe(swipeModel)
	if err != nil {

		// check if the error is due to a duplicate swipe
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				types.ErrResp(c, http.StatusConflict, "swipe already exists", nil)
				return
			}
		}

		types.ErrResp(c, http.StatusInternalServerError, "error creating swipe", nil)
		return
	}

	// check if there is a match
	rightSwipe, err := dao.CheckSwipeRight(swipeModel.To, swipeModel.From)
	if err != nil {
		types.ErrResp(c, http.StatusInternalServerError, "error checking swipe", nil)
		return
	}

	if rightSwipe {
		// create a match record as well
		match := types.Match{
			UserOneID: swipeModel.From,
			UserTwoID: swipeModel.To,
		}

		_, err := dao.CreateMatch(match)
		if err != nil {
			types.ErrResp(c, http.StatusInternalServerError, "error creating match", nil)
			return
		}

		types.OkResp(c, http.StatusOK, types.SwipeResponse{Matched: true, MatchId: &match.UserTwoID})
	} else {
		types.OkResp(c, http.StatusOK, types.SwipeResponse{Matched: false, MatchId: nil})
	}
}
