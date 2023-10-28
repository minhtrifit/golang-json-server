package controllers

import (
	"github.com/gin-gonic/gin"
)

func HandleRunServer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Server run succesfully",
	})
}