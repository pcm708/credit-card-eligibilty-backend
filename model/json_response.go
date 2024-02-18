package model

type JSONResponse struct {
	Status string `json:"status"`
}

type JSONFor400 struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
