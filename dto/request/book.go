package request

type CreateBookDto struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"required"`
}

type UpdateBookDto struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   int    `json:"year" validate:"required"`
}
