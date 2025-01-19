package main

import (
	"net/http" // Package http provides HTTP client and server implementations.

	"github.com/go-chi/chi/v5"            // Package chi is a lightweight, idiomatic router for building Go HTTP services.
	"github.com/go-chi/chi/v5/middleware" // Package middleware provides reusable middleware for Chi, like Heartbeat.
	"github.com/go-chi/cors"              // Package cors provides middleware to handle Cross-Origin Resource Sharing (CORS).
)

// routes defines the HTTP routes for the broker service.
// It configures middleware for CORS (Cross-Origin Resource Sharing), Heartbeat, and the route for the broker.
func (app *Config) routes() http.Handler {
	// Create a new Chi router for handling HTTP routes.
	mux := chi.NewRouter()

	// CORS middleware setup: Specify who is allowed to connect to the server.
	// This configuration allows any domain (http and https), with specific allowed methods and headers.
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},                                   // Allow all subdomains of http and https.
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 // Define allowed HTTP methods.
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Specify the allowed HTTP headers.
		ExposedHeaders:   []string{"Link"},                                                    // Expose the 'Link' header to browsers.
		AllowCredentials: true,                                                                // Allow credentials (cookies, authentication).
		MaxAge:           300,                                                                 // Cache preflight response for 5 minutes to optimize requests.
	}))

	// Heartbeat middleware to handle health check requests at "/ping" URL.
	// It helps monitor whether the service is alive and accessible.
	mux.Use(middleware.Heartbeat("/ping"))

	// Define the main POST route for the broker service at the root URL ("/").
	// It connects to the Broker method to handle the request.
	mux.Post("/", app.Broker)

	// Return the configured Chi router as the HTTP handler.
	return mux
}
