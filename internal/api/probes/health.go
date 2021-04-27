package probes

import (
	"encoding/json"
	"log"
	"net/http"
)

type status struct {
	Code string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	status := status{Code: "OK"}

	body, err := json.Marshal(status)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
