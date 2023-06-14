package main

import (
	"log"

	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/spf13/viper"
)

func init() {
	err := initializers.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	initializers.ConnectToDB()
	// scripts.SeedDummyData()
	// initializers.AutoMigrate()
}

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: helpers.ErrorHandler,
		BodyLimit:    10 * 1024 * 1024,
	})

	app.Use(helmet.New())
	// app.Use(logger.New(logger.Config{}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))
	// app.Use(limiter.New())

	app.Static("/", "./public")

	routers.Config(app)

	app.Listen(":" + viper.GetString("PORT"))

}
