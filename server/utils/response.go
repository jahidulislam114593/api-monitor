package utils

import (
	"encoding/json"
	"net/http"

	"github.com/api-monitor/models"
)

// SendJSONResponse sends a JSON response
func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// SendSuccessResponse sends a standardized success response
func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	SendJSONResponse(w, response)
}

// SendErrorResponse sends a standardized error response
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	response := models.APIResponse{
		Success: false,
		Error:   message,
	}
	SendJSONResponse(w, response)
}

// SendValidationError sends a 400 Bad Request with validation error
func SendValidationError(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusBadRequest, message)
}

// SendNotFoundError sends a 404 Not Found error
func SendNotFoundError(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusNotFound, message)
}

// SendInternalServerError sends a 500 Internal Server Error
func SendInternalServerError(w http.ResponseWriter, message string) {
	SendErrorResponse(w, http.StatusInternalServerError, message)
}