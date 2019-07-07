package main

import (
	"log"
	"github.com/philaporter/random/app"
	"syscall"
)

func main() {
	e := make(chan error)
	go app.StartServer(e)
	for {
		select {
		case err := <-e:
			log.Println("Error received", err)
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}
	}
}