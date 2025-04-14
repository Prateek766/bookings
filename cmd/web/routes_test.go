package main

import (
	"net/http"
	"testing"

	"github.com/Prateek766/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case http.Handler:
		// Do nothing
	default:
		t.Errorf("type is not http.Handler, but it is %T", v)
	}

}
