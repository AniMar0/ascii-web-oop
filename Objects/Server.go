package ascii_web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var (
	Input    Inputs
	stOutput Output
	Err      Error
)

// Define the Server struct
type Server struct{}

// Declare a variable to hold a template, initialized as a pointer
var temp *template.Template

// Run method starts the HTTP server and sets up routes
func (ser Server) Run() {
	// Set up a route for the each root URL  that calls the correspending method
	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/proccese", procceseHandler)
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/error", errorHandler)
	// Start the HTTP server on port 8080 and log any panic that occurs
	// If the server encounters an error, it will panic and terminate the program
	panic(http.ListenAndServe(":8080", nil))
}

func staticHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		homeHandler(res, req)
		return
	}
	path := "./templates/static/" + req.URL.Path

	fileData, err := os.ReadFile(path)
	if err != nil {
		http.Error(res, "404 - page not found", 404)
		return
	}
	fileName := strings.Split(path, ".")
	ext := fileName[len(fileName)-1]
	switch ext {
	case "html":
		res.Header().Set("Content-Type", "text/html")
	case "css":
		res.Header().Set("Content-Type", "text/css")
	default:
		res.Header().Set("Content-Type", "text/plain")
	}

	res.Write(fileData)
}

// homeHandler handles requests to the home page ("/")
func homeHandler(writer http.ResponseWriter, request *http.Request) {
	// Reset any previous errors using the custom error handling method
	stOutput.ResetOutput()
	Input.ResetInput()
	Err.ResetError()
	if request.URL.Path != "/" {
		// Render a 404 error page if the path is not found
		Err.RenderErrorPage(writer, fmt.Errorf("page not found"), 404) // 404 is status
		return
	}
	// Check if the request method is GET
	if request.Method == http.MethodGet {
		temp, Err.webError = template.ParseFiles("templates/index.html")
		if Err.webError != nil {
			Err.ErrGen(writer)
			return
		}
		temp.Execute(writer, stOutput)
	} else {
		Err.methodError = fmt.Errorf("request Err")
		// Execute the parsed template and write the output to the response
		Err.ErrGen(writer)
	}
}

// errorHandler handles requests to the error page
func errorHandler(writer http.ResponseWriter, request *http.Request) {
	// Reset any previous errors using the custom error handling method
	Err.ResetError()

	temp, Err.webError = template.ParseFiles("templates/error.html")
	if Err.webError != nil {
		Err.ErrGen(writer)
		return
	}
	// Execute the parsed error template and write the output to the response
	temp.Execute(writer, Err)
}

// procceseHandler handles requests for processing input and generating output
func procceseHandler(writer http.ResponseWriter, request *http.Request) {
	// Reset the output state, any input values and any previous errors
	stOutput.ResetOutput()
	Input.ResetInput()
	Err.ResetError()

	if request.Method == http.MethodPost {
		temp, Err.webError = template.ParseFiles("templates/index.html")
		if Err.webError != nil {
			Err.ErrGen(writer)
			return
		}
		request.ParseForm()

		// Retrieve values from the form and assign them to Input fields
		Input.UsreText = request.FormValue("input")
		Input.UserBanner = request.FormValue("banner")
		// Read a file based on the user input
		Input.ReadFile()

		if Err.bannerError != nil {
			Err.ErrGen(writer)
			return
		}

		// Generate art based on the user input
		Input.GenTheArt()
		if Err.inputError != nil {
			Err.ErrGen(writer)
			return
		}

		// Prepare the output by appending the user text
		stOutput.Old = "\n" + Input.UsreText
		// Generate the final output
		stOutput.GenAll(Input)
		// Execute the parsed template and write the output to the response
		temp.Execute(writer, stOutput)

	} else {
		// If the request method is not POST, generate an error
		Err.methodError = fmt.Errorf("request Err")
		Err.ErrGen(writer)
	}
}

func downloadHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		stOutput.Download(writer)
	} else {
		Err.methodError = fmt.Errorf("request Err")
		Err.ErrGen(writer)
	}
}
