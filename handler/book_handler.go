package handler

import (
	"crud_fiber.com/m/dto/request"
	"crud_fiber.com/m/pkg/errs"
	"crud_fiber.com/m/service"
	"crud_fiber.com/m/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		return utils.ErrorResponse(c, &errs.BadRequestError{
			Err: err.Error(),
		})
	}

	_, err := handler.BookService.CreateBook(payload)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, "Book created successfully")
}

// GetBooks @Summary Get all books
// @Description Get a list of books
// @Tags Books
// @Security OAuth2Password
// @Accept json
// @Produce json
// @Success 200 {object} documentation.ListBookResponse
// @Router /api/v1/book/list [get]
func (handler *BookHandler) GetBooks(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	books, err := handler.BookService.GetListPaginationBooks(page, limit)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, books)
}

func (handler *BookHandler) GetDetailBook(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return utils.ErrorResponse(c, errParse)
	}

	book, err := handler.BookService.GetDetailBook(id)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, book)
}

func (handler *BookHandler) UpdateBook(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return utils.ErrorResponse(c, errParse)
	}

	var payload request.UpdateBookDto
	if err := c.BodyParser(&payload); err != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{
			Err: err.Error(),
		})
	}

	_, err := handler.BookService.UpdateBook(id, payload)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, "Book updated successfully")
}

func (handler *BookHandler) DeleteBook(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, errParse := uuid.Parse(idStr)
	if errParse != nil {
		return utils.ErrorResponse(c, errParse)
	}

	_, err := handler.BookService.DeleteBook(id)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, "Book deleted successfully")
}
