package database

import "github.com/api-monitor/models"

// In-memory storage for API endpoints and checks
var (
	APIEndpoints []models.APIEndpoint
	APIChecks    []models.APICheck
)

// Initialize with some sample data for demo
func init() {
	// Sample endpoints for demonstration
	APIEndpoints = []models.APIEndpoint{
		{
			ID:          1,
			Name:        "GitHub API",
			URL:         "https://api.github.com/users/octocat",
			Method:      "GET",
			Description: "GitHub user endpoint",
			IsActive:    true,
		},
		{
			ID:          2,
			Name:        "JSONPlaceholder",
			URL:         "https://jsonplaceholder.typicode.com/posts/1",
			Method:      "GET", 
			Description: "Sample JSON API for testing",
			IsActive:    true,
		},
		{
			ID:          3,
			Name:        "HTTPBin",
			URL:         "https://httpbin.org/get",
			Method:      "GET",
			Description: "HTTP testing service",
			IsActive:    true,
		},
		{
			ID:          4,
			Name:        "Local Server",
			URL:         "http://localhost:3000/api/health",
			Method:      "GET",
			Description: "Local development server",
			IsActive:    false,
		},
	}

	// Sample checks data (some initial data for demo)
	APIChecks = []models.APICheck{
		{
			ID:           1,
			URL:          "https://api.github.com/users/octocat",
			ResponseTime: 245,
			Status:       200,
			CheckedAt:    "2024-01-15 14:30:25",
			IsUp:         true,
		},
		{
			ID:           2,
			URL:          "https://jsonplaceholder.typicode.com/posts/1",
			ResponseTime: 156,
			Status:       200,
			CheckedAt:    "2024-01-15 14:30:30",
			IsUp:         true,
		},
		{
			ID:           3,
			URL:          "https://httpbin.org/get",
			ResponseTime: 320,
			Status:       200,
			CheckedAt:    "2024-01-15 14:30:35",
			IsUp:         true,
		},
	}
}


func GetNextEndpointID() int {
	maxID := 0
	for _, endpoint := range APIEndpoints {
		if endpoint.ID > maxID {
			maxID = endpoint.ID
		}
	}
	return maxID + 1
}

func GetNextCheckID() int {
	maxID := 0
	for _, check := range APIChecks {
		if check.ID > maxID {
			maxID = check.ID
		}
	}
	return maxID + 1
}

func FindEndpointByID(id int) *models.APIEndpoint {
	for i, endpoint := range APIEndpoints {
		if endpoint.ID == id {
			return &APIEndpoints[i]
		}
	}
	return nil
}

func DeleteEndpointByID(id int) bool {
	for i, endpoint := range APIEndpoints {
		if endpoint.ID == id {
			APIEndpoints = append(APIEndpoints[:i], APIEndpoints[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateEndpointByID(id int, updatedEndpoint models.APIEndpoint) bool {
	for i, endpoint := range APIEndpoints {
		if endpoint.ID == id {
			updatedEndpoint.ID = id
			APIEndpoints[i] = updatedEndpoint
			return true
		}
	}
	return false
}

func GetActiveEndpoints() []models.APIEndpoint {
	var activeEndpoints []models.APIEndpoint
	for _, endpoint := range APIEndpoints {
		if endpoint.IsActive {
			activeEndpoints = append(activeEndpoints, endpoint)
		}
	}
	return activeEndpoints
}