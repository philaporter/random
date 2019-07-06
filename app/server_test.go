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
		tick := time.NewTicker(time.Second * 16).C
		go Start(err)
		select {
		case <-err:
			t.Error("Error with the server")
		case <-tick:
			return
		}
	})
}

func TestToggleHealthBool(t *testing.T) {
	check.Store(false)
	t.Run("Test health check toggle", func(t *testing.T) {
		toggleHealthBool()
		if check.Load() != true {
			t.Error("Toggle should have flipped check to true")
		}
	})
	t.Run("Test health check toggle", func(t *testing.T) {
		toggleHealthBool()
		if check.Load() != false {
			t.Error("Toggle should have flipped check to false")
		}
	})
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
