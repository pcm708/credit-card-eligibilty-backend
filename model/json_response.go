package model

// JSONResponse represents the structure of the JSON response
type JSONResponse struct {
	Status string `json:"status"`
}

// JsonForError represents the structure of the JSON response for 4XX errors
type JsonError struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// LogEntry represents the structure of the log entry
type LogEntry struct {
	Request_ID string `json:"request-id"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
}
