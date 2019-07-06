package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

// TODO: Create a process to check how long the service has been unhealthy and then shutdown the app if it's been too long

// Start starts the HTTP server
func Start(err chan error) {

	// Register shutdown listener and pass in the toggle bool channel
	go StartShutdownListener()

	// create route for HEALTH
	r := mux.NewRouter()
	r.HandleFunc(HEALTH, HealthHandler).Methods(GET)

	go reallyStart(err, r)
	logger.Info("Server successfully started")
	toggleHealthBool()
}

func reallyStart(err chan error, r *mux.Router) {
	// start server
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
	logger.Debug("the service is healthy: ", check.Load())
}

// HealthHandler checks the HealthCheckBool to return 200 or 503
func HealthHandler(writer http.ResponseWriter, request *http.Request) {
	if check.Load().(bool) {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusServiceUnavailable)
	}
}
