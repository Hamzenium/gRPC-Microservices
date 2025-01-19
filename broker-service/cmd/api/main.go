package main

import (
	"fmt"      // Package fmt implements formatted I/O with functions similar to C's printf and scanf.
	"log"      // Package log implements simple logging.
	"net/http" // Package http provides HTTP client and server implementations.
)

// webPort defines the port on which the broker service will listen for incoming requests.
const webPort = "80"

// Config is a placeholder struct to represent the application's configuration.
// Currently, it doesn't hold any values but can be expanded in the future.
type Config struct{}

// main is the entry point of the application.
// It initializes the app, starts the HTTP server, and handles errors if any occur.
func main() {
	// Create an instance of Config (although empty in this case).
	app := Config{}

	// Log that the broker service is starting and the port it will use.
	log.Printf("Starting broker service on port %s\n", webPort)

	// Define the HTTP server with the specified port and the router/handler.
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort), // Specify the address and port to listen on (e.g., :80).
		Handler: app.routes(),                // Assign the router (which handles the routes for the app).
	}

	// Start the HTTP server using ListenAndServe method.
	// If there's an error, log it and terminate the application.
	err := srv.ListenAndServe()
	if err != nil {
		// Log the error and stop the application if it fails to start the server.
		log.Panic(err)
	}
}
