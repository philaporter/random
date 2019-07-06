package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartShutdownListener() {

	// Channel to listen for os.Signal
	sigs := make(chan os.Signal, 1)
	shutdown := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Listen for system interrupts, then send shutdown signal when found
	go func() {
		sig := <-sigs
		log.Println("Received", sig)
		shutdown <- true
	}()

	// Wait for system interrupt
	<-shutdown
	shutdownHandler()
}

func shutdownHandler() {
	log.Println("Shutdown process started")
	timer := time.NewTicker(time.Second * 5).C
	<-timer
	log.Println("Exiting application")
	os.Exit(1)
}
