package main

import (
	"backend/Router"

	//"backend/Service/DB"

	//"backend/Model/TestModel"

	"github.com/gin-gonic/gin"
)

func main() {
	//DB.Init()
	r := gin.Default()
	Router.MainRoute(r)
	//TestModel.InsirtTestModel("test", "test")
	//TestModel.GetAllTestModel()
	r.Run("0.0.0.0:1234")

}
