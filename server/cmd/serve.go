package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/api-monitor/handlers"
	"github.com/api-monitor/middleware"
)

func Serve() {
	mux := http.NewServeMux()
	
	// API Endpoints management routes
	mux.Handle("GET /api/endpoints", http.HandlerFunc(handlers.GetAPIEndpoints))
	mux.Handle("POST /api/endpoints", http.HandlerFunc(handlers.CreateAPIEndpoint))
	mux.Handle("PUT /api/endpoints/{id}", http.HandlerFunc(handlers.UpdateAPIEndpoint))
	mux.Handle("DELETE /api/endpoints/{id}", http.HandlerFunc(handlers.DeleteAPIEndpoint))

	// API Checks routes (the monitoring functionality)
	mux.Handle("GET /api/checks", http.HandlerFunc(handlers.GetAPIChecks))
	mux.Handle("POST /api/checks/run/{endpoint_id}", http.HandlerFunc(handlers.RunAPICheck))
	mux.Handle("POST /api/checks/run-all", http.HandlerFunc(handlers.RunAllAPIChecks))

	// Statistics and monitoring overview
	mux.Handle("GET /api/stats", http.HandlerFunc(handlers.GetMonitoringStats))

	// Health check for the monitoring service itself
	mux.Handle("GET /api/health", http.HandlerFunc(handlers.HealthCheck))

	log.Println("🚀 API Response Time Monitor Server starting on http://localhost:8080")
	log.Println("📊 Available endpoints:")
	log.Println("   GET  /api/endpoints     - List all monitored endpoints")
	log.Println("   POST /api/endpoints     - Add new endpoint to monitor")
	log.Println("   PUT  /api/endpoints/{id} - Update endpoint")
	log.Println("   DELETE /api/endpoints/{id} - Delete endpoint")
	log.Println("   GET  /api/checks        - View check history")
	log.Println("   POST /api/checks/run/{endpoint_id} - Run check on specific endpoint")
	log.Println("   POST /api/checks/run-all - Run checks on all active endpoints")
	log.Println("   GET  /api/stats         - View monitoring statistics")
	log.Println("   GET  /api/health        - Service health check")

	fmt.Println("Server is running at : 8080")
	if err := http.ListenAndServe(":8080", middleware.CorsMiddleware(mux)); err != nil {
		fmt.Println("Error while starting server", err)
		return
	}
}