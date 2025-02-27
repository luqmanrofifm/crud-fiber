package apps

import (
	"crud_fiber.com/m/config"
	"crud_fiber.com/m/config/database"
	"crud_fiber.com/m/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApps() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	app.Use(middleware.CustomRecoverMiddleware)

	config.LoadEnv()
	database.InitializeDatabase()

	errApp := app.Listen(":8080")
	if errApp != nil {
		fmt.Println("Error when running the app")
	}
}
