package main

import (
	"caffanap/bookstore/api/config"
	"caffanap/bookstore/api/routers"
)

func main() {
	conf := config.LoadConfig()
	config.ConnectDatabase(conf.UrlConnectionMysql)
	config.MigrationDatabase()

	app := conf.App
	defer app.Listen(conf.AppPort)

	routers.HandleRouter(app)
}
