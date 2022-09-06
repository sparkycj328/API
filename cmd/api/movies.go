package main 

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// createMovieHandler() will handle the POST request to the /v1/movies endpoint
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// showMovieHandler() will handle the GET request to the /v1/movies:id endpoint
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the URL string for the ID parameter and will store all parameter names and values
	// into a slice
	params := httprouter.ParamsFromContext(r.Context())

	// Use ByName method to retrieve value of the id Paramater from the slice stored in params
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}
	
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}