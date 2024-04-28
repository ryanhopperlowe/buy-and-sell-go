package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/controller"
	"github.com/ryanhopperlowe/buy-and-sell-go/db"
	"github.com/ryanhopperlowe/buy-and-sell-go/initializers"
)

func main() {
	initializers.Init()

	r := gin.Default()

	controller.MakeListingRoutes(r, db.DB)
	controller.MakeUserRoutes(r, db.DB)

	r.Run()
}
