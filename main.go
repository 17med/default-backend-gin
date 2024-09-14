package main

import (
	"backend/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Router.MainRoute(r)
	r.Run(":1234")

}