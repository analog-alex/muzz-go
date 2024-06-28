package controllers

import "github.com/gin-gonic/gin"

func Discover(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Match",
	})
}

func Swipe(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Swipe",
	})
}
