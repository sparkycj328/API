package main 

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sparkycj328/API/internal/data"
)

// createMovieHandler() will handle the POST request to the /v1/movies endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// showMovieHandler() will handle the GET request to the /v1/movies:id endpoint
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return 
	}

	// Create a new instance of the Movie struct
	movie := data.Movie{
		ID:			id,
		CreatedAt:   time.Now(),
		Title:		"Casablanca",
		Runtime: 	102,
		Genres:		[]string{"drama", "romance", "war"},
		Version:	1,
	}


	app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}
}