package model

type JSONResponse struct {
	Status string `json:"status"`
}

type JsonFor4XX struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type LogEntry struct {
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}
