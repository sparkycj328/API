package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() * httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Convert notFoundResonse() to an http.Handler using the http.HandlerFunc() adapter
	// and set it as customer error handler for 404 responses
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	// Do the same for methodnotAllowedResponse and set it as custom 405 error response
	router.MethodNotAllowed = http.HandlerFunc(app. methodNotAllowedResponse)

	// Register methods, URL patterns and handler functions
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}