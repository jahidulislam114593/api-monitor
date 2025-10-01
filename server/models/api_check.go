package models

import "time"

type APICheck struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`           // "https://api.github.com/users"
	ResponseTime int    `json:"response_time"` // milliseconds
	Status       int    `json:"status"`        // 200, 404, 500
	CheckedAt    string `json:"checked_at"`    // timestamp
	IsUp         bool   `json:"is_up"`
	ErrorMessage string `json:"error_message,omitempty"` // if request failed
}

// Helper function to create a timestamp
func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Create a new check instance
func NewAPICheck(url string) *APICheck {
	return &APICheck{
		URL:       url,
		CheckedAt: GetCurrentTimestamp(),
	}
}

// Mark check as failed
func (c *APICheck) MarkAsFailed(err error) {
	c.IsUp = false
	c.Status = 0
	c.ErrorMessage = err.Error()
}

// Mark check as successful
func (c *APICheck) MarkAsSuccessful(statusCode int, responseTimeMs int) {
	c.Status = statusCode
	c.ResponseTime = responseTimeMs
	c.IsUp = statusCode >= 200 && statusCode < 400
}