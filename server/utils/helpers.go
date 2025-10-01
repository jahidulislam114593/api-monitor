package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/api-monitor/database"
	"github.com/api-monitor/models"
)

// ParseIDFromPath extracts and validates ID from URL path
func ParseIDFromPath(r *http.Request, paramName string) (int, error) {
	idStr := r.PathValue(paramName)
	if idStr == "" {
		return 0, fmt.Errorf("missing %s parameter", paramName)
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid %s parameter", paramName)
	}
	
	return id, nil
}

// PerformHTTPCheck performs an HTTP check on a URL and returns the result
func PerformHTTPCheck(url string) models.APICheck {
	check := models.APICheck{
		ID:        database.GetNextCheckID(),
		URL:       url,
		CheckedAt: models.GetCurrentTimestamp(),
	}

	// Record start time
	start := time.Now()

	// Make the HTTP request with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	
	// Calculate response time
	responseTime := int(time.Since(start).Milliseconds())
	
	if err != nil {
		// Request failed
		check.MarkAsFailed(err)
	} else {
		// Request succeeded
		check.MarkAsSuccessful(resp.StatusCode, responseTime)
		resp.Body.Close()
	}

	return check
}

// CalculateMonitoringStats calculates statistics from current data
func CalculateMonitoringStats() models.MonitoringStats {
	totalChecks := len(database.APIChecks)
	upCount := 0
	totalResponseTime := 0
	activeEndpoints := 0

	// Count active endpoints
	for _, endpoint := range database.APIEndpoints {
		if endpoint.IsActive {
			activeEndpoints++
		}
	}

	// Calculate check statistics
	for _, check := range database.APIChecks {
		if check.IsUp {
			upCount++
		}
		totalResponseTime += check.ResponseTime
	}

	var avgResponseTime float64
	if totalChecks > 0 {
		avgResponseTime = float64(totalResponseTime) / float64(totalChecks)
	}

	var uptimePercentage float64
	if totalChecks > 0 {
		uptimePercentage = (float64(upCount) / float64(totalChecks)) * 100
	}

	return models.MonitoringStats{
		TotalEndpoints:   len(database.APIEndpoints),
		ActiveEndpoints:  activeEndpoints,
		TotalChecks:      totalChecks,
		SuccessfulChecks: upCount,
		FailedChecks:     totalChecks - upCount,
		AvgResponseTime:  fmt.Sprintf("%.2f ms", avgResponseTime),
		UptimePercentage: fmt.Sprintf("%.2f%%", uptimePercentage),
	}
}

// FilterChecksByURL returns checks for a specific URL
func FilterChecksByURL(url string) []models.APICheck {
	var filtered []models.APICheck
	for _, check := range database.APIChecks {
		if check.URL == url {
			filtered = append(filtered, check)
		}
	}
	return filtered
}

// LimitChecks returns the last N checks
func LimitChecks(checks []models.APICheck, limit int) []models.APICheck {
	if limit <= 0 || limit >= len(checks) {
		return checks
	}
	
	// Return last N checks
	start := len(checks) - limit
	if start < 0 {
		start = 0
	}
	
	return checks[start:]
}