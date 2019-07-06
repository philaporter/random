package main

import (
	"github.com/philaporter/random/app"
	"log"
	"syscall"
)

func main() {
	e := make(chan error)
	go app.Start(e)
	for {
		select {
		case err := <-e:
			log.Println("Error received", err)
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}
	}
}