package middleware

import (
	"context"
	"crud_fiber.com/m/dto/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

func CustomRecoverMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			panicMessage := fmt.Sprintf("%v", r)
			log.Printf("Recovered from panic: %v", panicMessage)

			// trace detail error go runtime
			stackTrace := debug.Stack()
			log.Printf("Panic problem: %v", string(stackTrace))

			c.Status(http.StatusInternalServerError).JSON(response.Error{
				StatusCode: http.StatusInternalServerError,
				Error:      "INTERNAL_SERVER_ERROR",
				Message:    panicMessage,
				//Data:       string(stackTrace),
			})
		}

		duration := 240 * time.Second
		ctx, cancel := context.WithTimeout(c.Context(), duration)
		defer cancel()
		c.SetUserContext(ctx)
	}()
	return c.Next()
}
