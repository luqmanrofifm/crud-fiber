package utils

import (
	"crud_fiber.com/m/dto/response"
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(
		response.Success{
			StatusCode: fiber.StatusOK,
			Message:    "Success",
			Data:       data,
		},
	)
}

func ErrorResponse(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(
		response.Error{
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
			Error:      "Internal Server Error",
		},
	)
}
