package handler

import (
	"crud_fiber.com/m/dto/request"
	"crud_fiber.com/m/dto/response"
	"crud_fiber.com/m/service"
	"crud_fiber.com/m/utils"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	BookService *service.BookService
}

func NewBookHandler(bookService *service.BookService) *BookHandler {
	return &BookHandler{BookService: bookService}
}

func (handler *BookHandler) CreateBook(c *fiber.Ctx) error {
	var payload request.CreateBookDto
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			response.Error{
				StatusCode: fiber.StatusBadRequest,
				Message:    "Failed to parse request",
				Error:      err.Error(),
			})
	}

	_, err := handler.BookService.CreateBook(payload)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, "Book created successfully")
}
