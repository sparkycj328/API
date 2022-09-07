pacakge main

import (
	"fmt"
	"net/http"
)

// logError serves as a generic helper for logging error messages
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// errorResponse serves as a generic helper for sending JSON-formatted error
// messages to the client with a given status code.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serveErrorResponse will be used when our application encounters an unexpected problem at runtime.
// logs detailed errors then uses errorResponse to send a 500 Internal Server Error status code
// and JSON response to the client
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) { 
	message := "the requested resource could not be found" 
	app.errorResponse(w, r, http.StatusNotFound, message) } 
	
// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed 
// status code and JSON response to the client. func (app *application) 
methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) { 
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method) 
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message) 
}
