package routes

import (
	"crud_fiber.com/m/handler"
	"github.com/gofiber/fiber/v2"
)

type AuthRoute struct {
	App         *fiber.App
	AuthHandler *handler.AuthHandler
}

func (app *AuthRoute) SetupAuthRoute() {
	auth := app.App.Group("/api/v1/auth")

	auth.Post("/login", app.AuthHandler.Login)
	auth.Post("/register", app.AuthHandler.Register)
	auth.Post("/token", app.AuthHandler.GetOAuthToken)
	//auth.Post("/logout", app.AuthHandler.Logout)
	//auth.Post("/refresh", app.AuthHandler.Refresh)
	//auth.Get("/profile", app.AuthHandler.Profile)
	//auth.Post("/update-profile", app.AuthHandler.UpdateProfile)
	//auth.Post("/change-password", app.AuthHandler.ChangePassword)
	//auth.Post("/forgot-password", app.AuthHandler.ForgotPassword)
	//auth.Post("/reset-password", app.AuthHandler.ResetPassword)
}
