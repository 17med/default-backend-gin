package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {






	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}