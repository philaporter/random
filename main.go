package main

import (
	"github.com/alexcesaro/log/stdlog"
	"shutdown/app"
)

func main() {
	e := make(chan error)
	log := stdlog.GetFromFlags()
	go app.Start(e)

	for {
		select {
		case err := <-e:
			log.Error("An error was received on the error channel", err)
		}
	}
}

//t := time.NewTicker(time.Second * 30).C
//syscall.Kill(syscall.Getpid(), syscall.SIGINT)
//	case <-t:
//		go app.StartDuplicate(e)
//	}

// set health to 200, then toggle health to 503 after waiting 15 seconds
//go func() {
//	check.Store(true)
//	tick := time.NewTicker(time.Second * 15).C
//	<-tick
//	check.Store(!check.Load().(bool))
//}()

// Start starts the duplicate HTTP server using port already in use to trigger error
//func StartDuplicate(err chan error) {
//	oops := http.ListenAndServe(addressPort, nil)
//	if oops != nil {
//		err <- oops
//	}
//}
