package models

// Standard API response structure
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Statistics response structure
type MonitoringStats struct {
	TotalEndpoints     int     `json:"total_endpoints"`
	ActiveEndpoints    int     `json:"active_endpoints"`
	TotalChecks        int     `json:"total_checks"`
	SuccessfulChecks   int     `json:"successful_checks"`
	FailedChecks       int     `json:"failed_checks"`
	AvgResponseTime    string  `json:"avg_response_time"`
	UptimePercentage   string  `json:"uptime_percentage"`
}

// Bulk check response
type BulkCheckResponse struct {
	Message string     `json:"message"`
	Results []APICheck `json:"results"`
}