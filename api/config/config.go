package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Config struct {
	App                *fiber.App
	AppPort            string
	UrlConnectionMysql string
}

func LoadConfig() (config Config) {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	config.App = fiber.New()
	config.AppPort = os.Getenv(PORT)
	config.UrlConnectionMysql = mysqlUrlString()
	return config
}

func mysqlUrlString() string {
	db_host := os.Getenv(DATABASE_HOST)
	db_port := os.Getenv(DATABASE_PORT)
	db_username := os.Getenv(DATABASE_USERNAME)
	db_password := os.Getenv(DATABASE_PASSWORD)
	db_name := os.Getenv(DATABASE_NAME)

	return db_username + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name
}
