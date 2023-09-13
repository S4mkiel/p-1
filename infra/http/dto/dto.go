package dto

type APIResponseTrue struct {
	Success bool   `json:"success" default:"true"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type APIResponseFalse struct {
	Success bool  `json:"success" default:"false"`
	Error   []any `json:"error"`
}
