package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/listing"
	"github.com/ryanhopperlowe/buy-and-sell-go/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})

	db.AutoMigrate(&listing.Listing{})
	listing.MakeRoutes(r, db)

	db.AutoMigrate(&user.User{})
	user.MakeRoutes(r, db)

	r.Run()
}
