package Router

import (
	"backend/Controller/test"

	"backend/Middleware/testMid"

	"github.com/gin-gonic/gin"
)
func testR(rg *gin.RouterGroup) {
	rg.Use(testMid.SomeoneEnter)
	rg.GET("/",test.Test)
}