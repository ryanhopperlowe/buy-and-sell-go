package initializers

import "github.com/ryanhopperlowe/buy-and-sell-go/db"

func Init() {
	LoadEnv()
	db.InitDB()
	db.MigrateDB()
}
