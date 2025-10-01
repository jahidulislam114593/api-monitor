package models

type APIEndpoint struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`        // "GitHub API"
	URL         string `json:"url"`         // "https://api.github.com/users"
	Method      string `json:"method"`      // "GET", "POST"
	Description string `json:"description"` // "GitHub users endpoint"
	IsActive    bool   `json:"is_active"`   // whether to monitor this endpoint
}

// Validation methods
func (e *APIEndpoint) IsValid() bool {
	return e.Name != "" && e.URL != ""
}

func (e *APIEndpoint) SetDefaults() {
	if e.Method == "" {
		e.Method = "GET"
	}
	if e.ID == 0 {
		e.IsActive = true // default to active for new endpoints
	}
}