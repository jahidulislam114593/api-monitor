package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/api-monitor/database"
	"github.com/api-monitor/models"
	"github.com/api-monitor/utils"
)

// GetAPIChecks returns API check history with optional filtering
func GetAPIChecks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAPIChecks: Fetching check history")
	
	// Get query parameters for filtering
	urlFilter := r.URL.Query().Get("url")
	limitStr := r.URL.Query().Get("limit")
	
	var filteredChecks []models.APICheck
	
	// Filter by URL if provided
	if urlFilter != "" {
		log.Printf("GetAPIChecks: Filtering by URL: %s", urlFilter)
		filteredChecks = utils.FilterChecksByURL(urlFilter)
	} else {
		filteredChecks = database.APIChecks
	}
	
	// Apply limit if provided
	if limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			log.Printf("GetAPIChecks: Applying limit: %d", limit)
			filteredChecks = utils.LimitChecks(filteredChecks, limit)
		}
	}
	
	log.Printf("GetAPIChecks: Returning %d checks", len(filteredChecks))
	utils.SendJSONResponse(w, filteredChecks)
}

// RunAPICheck runs a health check on a specific endpoint
func RunAPICheck(w http.ResponseWriter, r *http.Request) {
	endpointID, err := utils.ParseIDFromPath(r, "endpoint_id")
	if err != nil {
		log.Printf("RunAPICheck: %v", err)
		utils.SendValidationError(w, "Invalid endpoint ID")
		return
	}

	log.Printf("RunAPICheck: Running check for endpoint ID: %d", endpointID)

	// Find the endpoint
	targetEndpoint := database.FindEndpointByID(endpointID)
	if targetEndpoint == nil {
		log.Printf("RunAPICheck: Endpoint with ID %d not found", endpointID)
		utils.SendNotFoundError(w, "Endpoint not found")
		return
	}

	log.Printf("RunAPICheck: Found endpoint %s (%s)", targetEndpoint.Name, targetEndpoint.URL)

	// Perform the actual API check
	checkResult := utils.PerformHTTPCheck(targetEndpoint.URL)
	
	// Save the check result
	database.APIChecks = append(database.APIChecks, checkResult)

	log.Printf("RunAPICheck: Check completed - Status: %d, Response Time: %dms, Up: %t", 
		checkResult.Status, checkResult.ResponseTime, checkResult.IsUp)

	utils.SendJSONResponse(w, checkResult)
}

// RunAllAPIChecks runs health checks on all active endpoints
func RunAllAPIChecks(w http.ResponseWriter, r *http.Request) {
	log.Println("RunAllAPIChecks: Starting bulk check operation")
	
	activeEndpoints := database.GetActiveEndpoints()
	log.Printf("RunAllAPIChecks: Found %d active endpoints", len(activeEndpoints))

	if len(activeEndpoints) == 0 {
		log.Println("RunAllAPIChecks: No active endpoints to check")
		response := models.BulkCheckResponse{
			Message: "No active endpoints to check",
			Results: []models.APICheck{},
		}
		utils.SendJSONResponse(w, response)
		return
	}

	var results []models.APICheck

	for _, endpoint := range activeEndpoints {
		log.Printf("RunAllAPIChecks: Checking %s (%s)", endpoint.Name, endpoint.URL)
		
		checkResult := utils.PerformHTTPCheck(endpoint.URL)
		database.APIChecks = append(database.APIChecks, checkResult)
		results = append(results, checkResult)
		
		log.Printf("RunAllAPIChecks: %s - Status: %d, Response: %dms", 
			endpoint.Name, checkResult.Status, checkResult.ResponseTime)
	}

	log.Printf("RunAllAPIChecks: Completed checks on %d endpoints", len(results))

	response := models.BulkCheckResponse{
		Message: fmt.Sprintf("Ran checks for %d active endpoints", len(results)),
		Results: results,
	}

	utils.SendJSONResponse(w, response)
}

// GetMonitoringStats returns comprehensive monitoring statistics
func GetMonitoringStats(w http.ResponseWriter, r *http.Request) {
	log.Println("GetMonitoringStats: Calculating monitoring statistics")
	
	stats := utils.CalculateMonitoringStats()
	
	log.Printf("GetMonitoringStats: Stats - Endpoints: %d, Checks: %d, Uptime: %s", 
		stats.TotalEndpoints, stats.TotalChecks, stats.UptimePercentage)

	utils.SendJSONResponse(w, stats)
}