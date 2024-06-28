package handlers

import (
	"github.com/gin-gonic/gin"
	"muzz-service/db"
	"net/http"
)

func Health(c *gin.Context) {
	err := db.GetDB().Ping(c)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "DB is not available",
		})
	}

	c.JSON(200, gin.H{
		"message": "OK",
	})
}
