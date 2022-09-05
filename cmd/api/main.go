package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config struct will hold all configuration setting for application.
type config struct {
	port int
	env string
}

// application struct will hold dependencies for the HTTP handlers, helpers, and middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	//declares an instance of the config struct
	var cfg config

	// Read in alues for port and env command-line flags into the config struct.
	// will default to number 4000 and environment if no flags provided
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which will write messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)


	//declare application struct which will store the newly defined config and logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servemux which will initiallly handle /v1/healthcheck route which dispatches requests
	// to the healthcheckHandler method
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// declare an HTTP server with reasonable timeout settings, will listen to port passed through the config struct
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}