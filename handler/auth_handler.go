package handler

import (
	"crud_fiber.com/m/dto/request"
	response_dto "crud_fiber.com/m/dto/response"
	"crud_fiber.com/m/pkg/errs"
	"crud_fiber.com/m/service"
	"crud_fiber.com/m/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (handler *AuthHandler) Register(c *fiber.Ctx) error {
	var payload request.RegisterDto
	if err := c.BodyParser(&payload); err != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{
			Err: err.Error(),
		})
	}

	err := handler.AuthService.Register(payload)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, "Register success")
}

func (handler *AuthHandler) Login(c *fiber.Ctx) error {
	var payload request.LoginDto
	if err := c.BodyParser(&payload); err != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{
			Err: err.Error(),
		})
	}

	token, err := handler.AuthService.Login(payload)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, token)
}

// GetOAuthToken Login @Summary Get OAuth Token
// @Description Mendapatkan token OAuth dengan username & password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.OAuthRequest true "OAuth Credentials"
// @Success 200 {object} response.OAuthResponse
// @Router /api/v1/auth/token [post]
func (handler *AuthHandler) GetOAuthToken(c *fiber.Ctx) error {
	var payload request.OAuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return utils.ErrorResponse(c, &errs.BadRequestError{
			Err: err.Error(),
		})
	}

	loginDto := request.LoginDto{
		Email:    payload.Username,
		Password: payload.Password,
	}
	token, err := handler.AuthService.Login(loginDto)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	response := response_dto.OAuthResponse{
		AccessToken: token.Token,
		TokenType:   "Bearer",
		ExpiresIn:   3600,
	}

	return c.JSON(response)
}
