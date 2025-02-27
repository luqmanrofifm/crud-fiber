package service

import (
	"crud_fiber.com/m/dto/request"
	"crud_fiber.com/m/dto/response"
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

func (service *BookService) GetListPaginationBooks(page int, limit int) (*response.PageListData, error) {
	books, totalData, err := service.BookRepository.FindAllPagination(page, limit)
	if err != nil {
		return nil, err
	}

	totalPage := (totalData + int64(limit) - 1) / int64(limit)

	return &response.PageListData{
		PageSize:    limit,
		CurrentPage: page,
		TotalPage:   totalPage,
		TotalRecord: totalData,
		Data:        books,
	}, nil
}
