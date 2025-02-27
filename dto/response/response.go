package response

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
)
