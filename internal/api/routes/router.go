package routes

import (
	"github.com/adalekin/otusk8s2/internal/api/probes"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health/", probes.HealthHandler).Methods("GET")
	return r
}
