package repository

import (
	"crud_fiber.com/m/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) Save(book *entity.Book) (*entity.Book, error) {
	err := r.db.Create(book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}
