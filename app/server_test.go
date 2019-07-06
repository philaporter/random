package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	err := make(chan error)
	t.Run("Test that the server will start", func(t *testing.T) {
		tick := time.NewTicker(time.Second * 1).C
		go Start(err)
		select {
		case <-err:
			t.Error("Error with the server")
		case <-tick:
			return
		}
	})
	//t.Run("Test that the duplicate server causes an error for the main server", func(t *testing.T) {
	//	tick := time.NewTicker(time.Second * 1).C
	//	go StartDuplicate(err)
	//	check := false
	//	for {
	//		select {
	//		case <-err:
	//			return
	//		case <-tick:
	//			go Start(err)
	//			if check {
	//				t.Error("The error path should have been triggered")
	//			}
	//			check = true
	//		}
	//	}
	//})
	//t.Run("Test that the server causes the duplicate server's error path", func(t *testing.T) {
	//	tick := time.NewTicker(time.Second * 1).C
	//	go Start(err)
	//	check := false
	//	for {
	//		select {
	//		case <-err:
	//			return
	//		case <-tick:
	//			go StartDuplicate(err)
	//			if check {
	//				t.Error("The error path should have been triggered")
	//			}
	//			check = true
	//		}
	//	}
	//})
}

func TestHealthHandler(t *testing.T) {
	t.Run("Test the isHealthy handler healthy path", func(t *testing.T) {
		check.Store(true)
		req, err := http.NewRequest(GET, HEALTH, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		check.Store(check.Load().(bool))
	})
	t.Run("Test the isHealthy handler unhealthy path", func(t *testing.T) {
		check.Store(false)
		req, err := http.NewRequest(GET, HEALTH, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusServiceUnavailable {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusServiceUnavailable)
		}
	})
}
