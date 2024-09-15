package main

import (
	"backend/Router"

	"backend/Service/DB"

	"backend/Model/TestModel"

	"github.com/gin-gonic/gin"
)

func main() {
	DB.Init()
	r := gin.Default()
	Router.MainRoute(r)
	TestModel.InsirtTestModel("test","test")
	r.Run(":1234")
	

}