package types

import "github.com/gin-gonic/gin"

type Extras map[string]any

type ErrorResponse struct {
	Message string  `json:"message"`
	Code    int     `json:"code"`
	Extras  *Extras `json:"extras,omitempty"`
}

func OkResp(c *gin.Context, code int, body any) {
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.JSON(code, body)
}

func ErrResp(c *gin.Context, code int, message string, extras *Extras) {
	errorResponse := &ErrorResponse{
		Message: message,
		Code:    code,
		Extras:  extras,
	}
	c.Writer.Header().Set("Backoff-Seconds", "0")
	c.JSON(code, errorResponse)
}
