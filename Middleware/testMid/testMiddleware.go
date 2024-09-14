package testMid

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func SomeoneEnter(ctx *gin.Context) {
	fmt.Println("someone enter test")
	ctx.Next()
}