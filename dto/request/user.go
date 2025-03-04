package request

type (
	RegisterDto struct {
		Email    string `json:"email" binding:"required,email"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	LoginDto struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	OAuthRequest struct {
		Username string `json:"username" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
)
