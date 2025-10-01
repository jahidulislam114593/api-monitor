package handlers

import (
	"net/http"
	"time"

	"github.com/api-monitor/database"
	"github.com/api-monitor/utils"
)

var startTime = time.Now()

// HealthCheck provides a health check endpoint for the service
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)
	
	response := map[string]interface{}{
		"status":      "healthy",
		"service":     "API Response Time Monitor",
		"version":     "1.0.0",
		"timestamp":   time.Now().Format(time.RFC3339),
		"uptime":      uptime.String(),
		"endpoints":   len(database.APIEndpoints),
		"checks":      len(database.APIChecks),
	}
	
	utils.SendJSONResponse(w, response)
}