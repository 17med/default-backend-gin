package Router

import (
	"github.com/gin-gonic/gin"
	

)
func MainRoute(r *gin.Engine) {
	testRouteA := r.Group("/test")
	
	testR(testRouteA)
	
}