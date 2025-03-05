package documentation

import "crud_fiber.com/m/entity"

type ListBookResponse struct {
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message"`
	Data       PageListBook `json:"data"`
}

type PageListBook struct {
	PageSize    int           `json:"page_size"`
	CurrentPage int           `json:"current_page"`
	TotalPage   int64         `json:"total_page"`
	TotalRecord int64         `json:"total_record"`
	Data        []entity.Book `json:"data"`
}
