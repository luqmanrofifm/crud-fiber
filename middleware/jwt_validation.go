package middleware

import (
	"crud_fiber.com/m/entity"
	"crud_fiber.com/m/pkg/errs"
	"crud_fiber.com/m/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type JwtValidation struct {
	db *gorm.DB
}

func NewJwtValidation(db *gorm.DB) *JwtValidation {
	return &JwtValidation{db}
}

func (j *JwtValidation) ValidateToken(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return utils.ErrorResponse(c, &errs.UnauthorizedError{
			Err: "Token Invalid",
		})
	}

	u := entity.User{}
	errValidateToken := u.ValidateToken(token)
	if errValidateToken != nil {
		return utils.ErrorResponse(c, errValidateToken)
	}

	user, err := j.FetchUserByEmail(u.Email)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	c.Locals("user", user)

	return c.Next()
}

func (j *JwtValidation) FetchUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := j.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
