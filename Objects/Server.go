package ascii_web

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	Input    Inputs
	stOutput Output
	Err      Error
)

type Server struct{}

var temp *template.Template

func (ser Server) Run() {
	fmt.Println("http://localhost:8080/")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/proccese", procceseHandler)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	Err.ResitError()
	if request.Method == http.MethodGet {
		temp, Err.webError = template.ParseFiles("templates/index.html")
		if Err.webError != nil {
			Err.ErrGen(writer)
			return
		}
		temp.Execute(writer, nil)
	} else {
		Err.methodError = fmt.Errorf("request Err")
		Err.ErrGen(writer)
	}
}

func errorHandler(writer http.ResponseWriter, request *http.Request) {
	Err.ResitError()
	temp, Err.webError = template.ParseFiles("templates/error.html")
	if Err.webError != nil {
		Err.ErrGen(writer)
		return
	}
	temp.Execute(writer, Err)
}

func procceseHandler(writer http.ResponseWriter, request *http.Request) {
	stOutput.ResitOutput()
	Input.ResitInput()
	Err.ResitError()

	if request.Method == http.MethodPost {
		temp, Err.webError = template.ParseFiles("templates/index.html")
		if Err.webError != nil {
			Err.ErrGen(writer)
			return
		}
		request.ParseForm()

		Input.UsreText = request.FormValue("input")
		Input.UserBanner = request.FormValue("banner")

		Input.ReadFile()

		if Err.bannerError != nil {
			Err.ErrGen(writer)
			return
		}
		Input.GenTheArt()
		if Err.inputError != nil {
			Err.ErrGen(writer)
			return
		}
		stOutput.GenAll(Input)

		temp.Execute(writer, stOutput)
	} else {
		fmt.Println("Method err")
		Err.methodError = fmt.Errorf("request Err")
		Err.ErrGen(writer)
	}
}
