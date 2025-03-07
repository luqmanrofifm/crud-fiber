package utils

import (
	"crud_fiber.com/m/dto/response"
	"crud_fiber.com/m/pkg/errs"
	"crud_fiber.com/m/pkg/logging"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"runtime/debug"
)

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(
		response.Success{
			StatusCode: fiber.StatusOK,
			Message:    "SUCCESS",
			Data:       data,
		},
	)
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	var badRequestError *errs.BadRequestError
	var resourceNotFoundError *errs.ResourceNotFoundError
	var unauthorizedError *errs.UnauthorizedError

	stackTrace := debug.Stack()

	logging.Log.Error(err.Error())

	if errors.As(err, &badRequestError) {
		return c.Status(http.StatusBadRequest).JSON(response.Error{
			StatusCode: http.StatusBadRequest,
			Error:      "BAD_REQUEST",
			Message:    err.Error(),
		})
	} else if errors.As(err, &resourceNotFoundError) {
		return c.Status(http.StatusNotFound).JSON(response.Error{
			StatusCode: http.StatusNotFound,
			Error:      "NOT_FOUND",
			Message:    err.Error(),
		})
	} else if errors.As(err, &unauthorizedError) {
		return c.Status(http.StatusUnauthorized).JSON(response.Error{
			StatusCode: http.StatusUnauthorized,
			Error:      "UNAUTHORIZED",
			Message:    err.Error(),
		})
	} else {
		log.Printf("error: %v", string(stackTrace))
		return c.Status(http.StatusInternalServerError).JSON(response.Error{
			StatusCode: http.StatusInternalServerError,
			Error:      "INTERNAL_SERVER_ERROR",
			Message:    err.Error(),
		})
	}
}
