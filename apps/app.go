package apps

import (
	"crud_fiber.com/m/config"
	"crud_fiber.com/m/config/database"
	"crud_fiber.com/m/handler"
	"crud_fiber.com/m/middleware"
	"crud_fiber.com/m/repository"
	"crud_fiber.com/m/routes"
	"crud_fiber.com/m/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
	"log"
)

func StartApps() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		//if c.Method() == "OPTIONS" {
		//	return c.SendStatus(fiber.StatusOK)
		//}
		return c.Next()
	})

	app.Use(middleware.CustomRecoverMiddleware)

	config.LoadEnv()
	database.InitializeDatabase()

	JwtValidation := middleware.NewJwtValidation(database.GetInstanceDatabase())

	bookRoute := routes.BookRoute{
		App:           app,
		BookHandler:   setUpBookHandler(database.GetInstanceDatabase()),
		JwtValidation: JwtValidation,
	}

	bookRoute.SetupBookRoute()

	authRoute := routes.AuthRoute{
		App:         app,
		AuthHandler: setUpAuthHandler(database.GetInstanceDatabase()),
	}

	authRoute.SetupAuthRoute()

	errApp := app.Listen(":8080")
	if errApp != nil {
		log.Println("Error when running the app")
	}
}

func setUpBookHandler(db *gorm.DB) *handler.BookHandler {
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	return bookHandler
}

func setUpAuthHandler(db *gorm.DB) *handler.AuthHandler {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)
	return authHandler
}
