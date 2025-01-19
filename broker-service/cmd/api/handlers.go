package main

import (
	"encoding/json" // Package for encoding and decoding JSON data.
	"net/http"      // Package for handling HTTP requests and responses.
)

// jsonResponse defines the structure of the response that will be returned from the Broker endpoint.
// It includes a flag for errors, a message, and optional data that may be included in the response.
type jsonResponse struct {
	Error   bool   `json:"error"`          // Indicates whether there was an error in the response.
	Message string `json:"message"`        // A message describing the result of the operation.
	Data    any    `json:"data,omitempty"` // Optional field to return additional data, it will be omitted if empty.
}

// Broker is the HTTP handler function for the broker endpoint.
// It sends a predefined JSON response when a request hits the broker route.
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Prepare the JSON response data, indicating no error and a message.
	payload := jsonResponse{
		Error:   false,            // No error occurred.
		Message: "Hit the broker", // Message describing the action.
	}

	// Marshal the payload into a pretty-printed JSON string (with indentation).
	out, _ := json.MarshalIndent(payload, "", "\t")

	// Set the response header to indicate that the content is JSON.
	w.Header().Set("Content-Type", "application/json")

	// Set the HTTP response status code to 202 Accepted.
	w.WriteHeader(http.StatusAccepted)

	// Write the JSON response to the response body.
	w.Write(out)
}
