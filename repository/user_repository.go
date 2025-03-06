package repository

import (
	"crud_fiber.com/m/entity"
	"crud_fiber.com/m/pkg/errs"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FetchUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, &errs.ResourceNotFoundError{Err: "User Not Found"}
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(u *entity.User) error {
	if err := r.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}
