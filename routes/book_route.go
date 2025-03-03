package routes

import (
	"crud_fiber.com/m/handler"
	"crud_fiber.com/m/middleware"
	"github.com/gofiber/fiber/v2"
)

type BookRoute struct {
	App           *fiber.App
	BookHandler   *handler.BookHandler
	JwtValidation *middleware.JwtValidation
}

func (route *BookRoute) SetupBookRoute() {
	bookRoute := route.App.Group("/api/v1/book", route.JwtValidation.ValidateToken)

	bookRoute.Post("/create", route.BookHandler.CreateBook)
	bookRoute.Get("/list", route.BookHandler.GetBooks)
	bookRoute.Get("/detail/:id", route.BookHandler.GetDetailBook)
	bookRoute.Put("/update/:id", route.BookHandler.UpdateBook)
	bookRoute.Delete("/delete/:id", route.BookHandler.DeleteBook)
}
