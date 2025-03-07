package handler

import (
	"crud_fiber.com/m/dto/request"
	"crud_fiber.com/m/entity"
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

// CreateBook @Summary Create a new book
// @Description Add a new book to the database
// @Tags Books
// @Security OAuth2Password
// @Accept json
// @Produce json
// @Param book body request.CreateBookDto true "Book data"
// @Success 200 {object} swagger.SuccessResponse
// @Router /api/v1/book/create [post]
func (handler *BookHandler) CreateBook(c *fiber.Ctx) error {
	var payload request.CreateBookDto
	if err := c.BodyParser(&payload); err != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{
			Err: err.Error(),
		})
	}

	if errValidate := utils.Validate(payload); errValidate != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{Err: errValidate.Error()})
	}

	user := c.Locals("user").(*entity.User)

	_, err := handler.BookService.CreateBook(payload, user)
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
// @Success 200 {object} swagger.ListBookResponse
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

// GetDetailBook @Summary Get detail book
// @Description Get detail of book
// @Tags Books
// @Security OAuth2Password
// @Accept json
// @Produce json
// @Success 200 {object} swagger.DetailBookResponse
// @Router /api/v1/book/detail/{id} [get]
// @Param id path string true "Book ID"
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

// UpdateBook @Summary Update a book
// @Description Update data of a book
// @Tags Books
// @Security OAuth2Password
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param payload body request.UpdateBookDto true "Book data"
// @Success 200 {object} swagger.SuccessResponse
// @Router /api/v1/book/update/{id} [put]
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

	if errValidate := utils.Validate(payload); errValidate != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{Err: errValidate.Error()})
	}

	user := c.Locals("user").(*entity.User)

	_, err := handler.BookService.UpdateBook(id, payload, user)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, "Book updated successfully")
}

// DeleteBook @Summary Delete a book
// @Description Delete a book from the database
// @Tags Books
// @Security OAuth2Password
// @Accept json
// @Produce json
// @Success 200 {object} swagger.SuccessResponse
// @Router /api/v1/book/delete/{id} [delete]
// @Param id path string true "Book ID"
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
