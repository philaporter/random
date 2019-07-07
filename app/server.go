package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// TODO: Create a process to check how long the service has been unhealthy and then random the app if it's been too long

// Start starts the HTTP server
func StartServer(err chan error) {

	log.Println("Beginning server startup - waiting 15 seconds to toggle /health as status OK")

	// Register random listener and pass in the toggle bool channel
	go startShutdownListener()

	// start the server for realz
	go start(err)

	// Wait 15 seconds to see if any errors are thrown, then set the service as healthy
	timer := time.NewTicker(time.Second * 15).C
	select {
	case <-timer:
		log.Println("Server successfully started - toggling /health to OK")
		toggleHealthBool()
	}
}

// reallyStart runs as a go routine and actually starts the http server
func start(err chan error) {

	// Create route for /health
	r := mux.NewRouter()
	r.HandleFunc(HEALTH, HealthHandler).Methods(GET)

	// Start server
	oops := http.ListenAndServe(ADDRESS_PORT, r)
	if oops != nil {
		err <- oops
	}
}

// Toggle the HEALTH bool for representing service HEALTH
func toggleHealthBool() {
	// set HEALTH bool to false as default
	if check.Load() == nil {
		check.Store(false)
	}

	// toggling HEALTH and logging the new status
	check.Store(!check.Load().(bool))
	log.Println("the service is healthy: ", check.Load())
}

// HealthHandler checks the HealthCheckBool to return 200 or 503
func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	if check.Load().(bool) {
		log.Println("Health check received: STATUS OK")
		writer.WriteHeader(http.StatusOK)
	} else {
		log.Println("Health check received: STATUS UNAVAILABLE")
		writer.WriteHeader(http.StatusServiceUnavailable)
	}
}
