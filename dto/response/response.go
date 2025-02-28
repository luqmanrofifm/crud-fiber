package response

import "github.com/google/uuid"

type (
	Success struct {
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data,omitempty"`
	}

	Error struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
		Error      string `json:"error,omitempty"`
	}

	PageListData struct {
		PageSize    int         `json:"page_size"`
		CurrentPage int         `json:"current_page"`
		TotalPage   int64       `json:"total_page"`
		TotalRecord int64       `json:"total_record"`
		Data        interface{} `json:"data"`
	}

	LoginResponse struct {
		Id    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Token string    `json:"token"`
	}
)
