package ascii_web

import (
	"fmt"
	"net/http"
	"text/template"
)

type Error struct {
	bannerError error
	inputError  error
	webError    error
	methodError error
	Err         string
	ErrNumber   string
}

func (err Error) ErrGen(w http.ResponseWriter) {
	switch {
	case err.bannerError != nil:
		if err.bannerError.Error() == "bad request" {
			err.RenderErrorPage(w, err.bannerError, 400)
		} else {
			err.RenderErrorPage(w, err.bannerError, 500)
		}

	case err.inputError != nil:
		err.RenderErrorPage(w, err.inputError, 500)
	case err.webError != nil:
		err.RenderErrorPage(w, fmt.Errorf("internal server error"), 500)
	case err.methodError != nil:
		err.RenderErrorPage(w, fmt.Errorf("method not allowed"), 405)
	}
}

func (err Error) RenderErrorPage(w http.ResponseWriter, errMsg error, errCode int) {
	tmpl, tempErr := template.ParseFiles("templates/error.html")
	if tempErr != nil {
		http.Error(w, tempErr.Error(), http.StatusNotFound)
		return
	}
	err = Error{Err: errMsg.Error(), ErrNumber: fmt.Sprintf("%d", errCode)}
	w.WriteHeader(errCode)
	tmpl.Execute(w, err)
}

func (err *Error) ResitError() {
	err.bannerError = nil
	err.inputError = nil
	err.methodError = nil
	err.webError = nil
	err.Err = ""
	err.ErrNumber = ""
}
