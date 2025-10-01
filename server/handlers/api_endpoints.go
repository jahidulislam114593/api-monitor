package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/api-monitor/database"
	"github.com/api-monitor/models"
	"github.com/api-monitor/utils"
)

// GetAPIEndpoints returns all API endpoints
func GetAPIEndpoints(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAPIEndpoints: Fetching %d endpoints", len(database.APIEndpoints))
	utils.SendJSONResponse(w, database.APIEndpoints)
}

// CreateAPIEndpoint creates a new API endpoint to monitor
func CreateAPIEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateAPIEndpoint: Starting endpoint creation")
	
	var newEndpoint models.APIEndpoint

	if err := json.NewDecoder(r.Body).Decode(&newEndpoint); err != nil {
		log.Printf("CreateAPIEndpoint: Error decoding endpoint data: %v", err)
		utils.SendValidationError(w, "Invalid JSON data")
		return
	}

	log.Printf("CreateAPIEndpoint: Decoded endpoint - Name: %s, URL: %s", newEndpoint.Name, newEndpoint.URL)

	// Validate the endpoint
	if !newEndpoint.IsValid() {
		log.Printf("CreateAPIEndpoint: Validation failed for endpoint: %+v", newEndpoint)
		utils.SendValidationError(w, "Name and URL are required")
		return
	}

	// Set defaults and generate ID
	newEndpoint.SetDefaults()
	newEndpoint.ID = database.GetNextEndpointID()
	
	// Add to database
	database.APIEndpoints = append(database.APIEndpoints, newEndpoint)

	log.Printf("CreateAPIEndpoint: Successfully created endpoint with ID: %d", newEndpoint.ID)
	log.Printf("CreateAPIEndpoint: Total endpoints in database: %d", len(database.APIEndpoints))

	utils.SendJSONResponse(w, newEndpoint)
}

// UpdateAPIEndpoint updates an existing API endpoint
func UpdateAPIEndpoint(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromPath(r, "id")
	if err != nil {
		log.Printf("UpdateAPIEndpoint: %v", err)
		utils.SendValidationError(w, "Invalid endpoint ID")
		return
	}

	log.Printf("UpdateAPIEndpoint: Attempting to update endpoint with ID: %d", id)

	var updateEndpoint models.APIEndpoint
	if err := json.NewDecoder(r.Body).Decode(&updateEndpoint); err != nil {
		log.Printf("UpdateAPIEndpoint: Error decoding endpoint data: %v", err)
		utils.SendValidationError(w, "Invalid JSON data")
		return
	}

	log.Printf("UpdateAPIEndpoint: Decoded update data - Name: %s, URL: %s", updateEndpoint.Name, updateEndpoint.URL)

	// Validate the endpoint
	if !updateEndpoint.IsValid() {
		log.Printf("UpdateAPIEndpoint: Validation failed for endpoint: %+v", updateEndpoint)
		utils.SendValidationError(w, "Name and URL are required")
		return
	}

	// Set defaults
	updateEndpoint.SetDefaults()

	// Update in database
	if database.UpdateEndpointByID(id, updateEndpoint) {
		log.Printf("UpdateAPIEndpoint: Successfully updated endpoint ID: %d", id)
		utils.SendJSONResponse(w, updateEndpoint)
	} else {
		log.Printf("UpdateAPIEndpoint: Endpoint with ID %d not found", id)
		utils.SendNotFoundError(w, "Endpoint not found")
	}
}

// DeleteAPIEndpoint deletes an API endpoint
func DeleteAPIEndpoint(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIDFromPath(r, "id")
	if err != nil {
		log.Printf("DeleteAPIEndpoint: %v", err)
		utils.SendValidationError(w, "Invalid endpoint ID")
		return
	}

	log.Printf("DeleteAPIEndpoint: Attempting to delete endpoint with ID: %d", id)

	if database.DeleteEndpointByID(id) {
		log.Printf("DeleteAPIEndpoint: Successfully deleted endpoint with ID: %d", id)
		log.Printf("DeleteAPIEndpoint: Remaining endpoints in database: %d", len(database.APIEndpoints))
		
		response := map[string]interface{}{
			"message": "Endpoint deleted successfully",
			"id":      id,
		}
		utils.SendJSONResponse(w, response)
	} else {
		log.Printf("DeleteAPIEndpoint: Endpoint with ID %d not found", id)
		utils.SendNotFoundError(w, "Endpoint not found")
	}
}