package ascii_web

import (
	"fmt"
	"net/http"
	"text/template"
)

// Error struct holds different types of errors that can occur
type Error struct {
	bannerError error  // Error related to banner file operations
	inputError  error  // Error related to user input
	webError    error  // General web-related errors
	methodError error  // Error related to unsupported HTTP methods
	Err         string // Custom error message
	ErrNumber   string // Custom error number (HTTP status code)
}

// ErrGen generates an appropriate error response based on the current state of errors
func (err Error) ErrGen(w http.ResponseWriter) {
	switch {
	case err.bannerError != nil:
		if err.bannerError.Error() == "bad request" {
			err.RenderErrorPage(w, err.bannerError, 400) // 400 Bad Request
		} else {
			err.RenderErrorPage(w, err.bannerError, 500) // 500 Internal Server Error
		}
	case err.inputError != nil:
		// If there is an input error, render the error page with a 400 status
		err.RenderErrorPage(w, err.inputError, 400)
	case err.webError != nil:
		// For general web errors, render a 500 error page
		err.RenderErrorPage(w, fmt.Errorf("internal server error"), 500)
	case err.methodError != nil:
		// For unsupported methods, render a 405 error page
		err.RenderErrorPage(w, fmt.Errorf("method not allowed"), 405)
	}
}

// RenderErrorPage renders the error page based on the provided error message and code
func (err Error) RenderErrorPage(w http.ResponseWriter, errMsg error, errCode int) {
	// Parse the error template file
	tmpl, tempErr := template.ParseFiles("templates/error.html")
	if tempErr != nil {
		http.Error(w, tempErr.Error(), http.StatusNotFound)
		return
	}
	// If template parsing fails, respond with an error
	// Create a new Error instance for rendering
	err = Error{Err: errMsg.Error(), ErrNumber: fmt.Sprintf("%d", errCode)}
	w.WriteHeader(errCode) // Set the HTTP status code
	tmpl.Execute(w, err)   // Render the template with the error data
}

// ResetError resets all error fields in the Error struct to their initial state
func (err *Error) ResetError() {
	err.bannerError = nil   // Clear banner error
	err.inputError = nil    // Clear input error
	err.methodError = nil   // Clear method error
	err.webError = nil      // Clear web error
	err.Err = ""            // Clear custom error message
	err.ErrNumber = ""      // Clear custom error number
}
