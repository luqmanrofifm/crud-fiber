package service

import (
	"crud_fiber.com/m/dto/request"
	"crud_fiber.com/m/entity"
	"crud_fiber.com/m/repository"
)

type BookService struct {
	BookRepository *repository.BookRepository
}

func NewBookService(bookRepository *repository.BookRepository) *BookService {
	return &BookService{BookRepository: bookRepository}
}

func (service *BookService) CreateBook(dto request.CreateBookDto) (*entity.Book, error) {
	book := entity.Book{
		Title:  dto.Title,
		Author: dto.Author,
		Year:   dto.Year,
	}

	createdBook, err := service.BookRepository.Save(&book)
	if err != nil {
		return nil, err
	}

	return createdBook, nil
}
