package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/controller"
	"github.com/ryanhopperlowe/buy-and-sell-go/initializers"
	"github.com/ryanhopperlowe/buy-and-sell-go/listing"
)

func Init() {
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.MigrateDB()
}

func main() {
	Init()

	r := gin.Default()

	listing.MakeRoutes(r, initializers.DB)
	controller.MakeUserRoutes(r, initializers.DB)

	r.Run()
}
