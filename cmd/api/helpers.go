package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"


	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" parameter from the current request context, then conver it
// to an integer and return it. If the operation isn't successful, return 0 and an error.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id paramater")
	}

	return id, nil	
}

// envelope type to help surround movie into a struct
type envelope map[string]interface{}

// writeJSON takes the destination http.ResponseWriter, the hTTP status code to send,
// data to encode to JSON, and a header map containing any aditional HTTP headers
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Append a newlike to the JSON
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	// add the JSON header and write the status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}