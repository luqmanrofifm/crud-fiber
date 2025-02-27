package routes

import (
	"crud_fiber.com/m/handler"
	"github.com/gofiber/fiber/v2"
)

type BookRoute struct {
	App         *fiber.App
	BookHandler *handler.BookHandler
}

func (route *BookRoute) SetupBookRoute() {
	bookRoute := route.App.Group("/api/v1/book")

	bookRoute.Post("/create", route.BookHandler.CreateBook)
	bookRoute.Get("/list", route.BookHandler.GetBooks)
	bookRoute.Get("/detail/:id", route.BookHandler.GetDetailBook)
}
