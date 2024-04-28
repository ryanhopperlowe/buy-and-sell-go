package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/controller"
	"github.com/ryanhopperlowe/buy-and-sell-go/initializers"
)

func Init() {
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.MigrateDB()
}

func main() {
	Init()

	r := gin.Default()

	controller.MakeListingRoutes(r, initializers.DB)
	controller.MakeUserRoutes(r, initializers.DB)

	r.Run()
}
