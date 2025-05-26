package main

import (
	"encoding/json"
	"net/http"
)

func handleStart(w http.ResponseWriter, r *http.Request) {
	session.StartTimer()
	respondJSON(w, session.State)
}

func handleStop(w http.ResponseWriter, r *http.Request) {
	session.StopTimer()
	respondJSON(w, session.State)
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, session.State)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	session.Reset()
	respondJSON(w, session.State)
}

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
