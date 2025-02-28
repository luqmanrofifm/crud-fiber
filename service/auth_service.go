package service

import (
	"crud_fiber.com/m/dto/request"
	"crud_fiber.com/m/entity"
	"crud_fiber.com/m/repository"
	"errors"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (a *AuthService) Register(dto request.RegisterDto) error {
	user := entity.User{
		CreatedBy: "SYSTEM",
		Email:     dto.Email,
		Password:  dto.Password,
		Name:      dto.Name,
	}

	user.HashPassword()

	errValidateEmail := a.validateDuplicateEmail(user.Email)
	if errValidateEmail != nil {
		return errValidateEmail
	}

	errCreateUser := a.UserRepository.CreateUser(&user)
	if errCreateUser != nil {
		return errCreateUser
	}

	return nil
}

func (a *AuthService) validateDuplicateEmail(email string) error {
	user, err := a.UserRepository.FetchUserByEmail(email)
	if err != nil && err.Error() != "record not found" {
		return err
	}

	if user != nil {
		return errors.New("email already registered")
	}

	return nil
}
