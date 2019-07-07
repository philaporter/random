# Random with Go

* health check handling
* graceful shutdown things
* http server with gorilla/mux
* dockerfile defined

## How to run

* `docker build -t random .`
* `docker run -p 8080:8080 random`
* GET to `localhost:8080/health` for health status

OR

* `./random`
* GET to `localhost:8080/health` for health status