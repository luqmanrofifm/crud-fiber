package swagger

import "github.com/google/uuid"

type SuccessResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       string `json:"data,omitempty"`
}

type LoginResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}
