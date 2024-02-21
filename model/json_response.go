package model

// JSONResponse represents the structure of the JSON response
type JSONResponse struct {
	Status string `json:"status"`
}

// JsonFor4XX represents the structure of the JSON response for 4XX errors
type JsonFor4XX struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// LogEntry represents the structure of the log entry
type LogEntry struct {
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Timestamp   string `json:"timestamp"`
}
