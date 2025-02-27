package repository

import (
	"crud_fiber.com/m/entity"
	"github.com/google/uuid"
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

func (r *BookRepository) FindAllPagination(page int, limit int) ([]entity.Book, int64, error) {
	var books []entity.Book
	offset := (page - 1) * limit

	err := r.db.Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	err = r.db.Model(&entity.Book{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return books, total, nil
}

func (r *BookRepository) FindByID(id uuid.UUID) (*entity.Book, error) {
	var book entity.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}
