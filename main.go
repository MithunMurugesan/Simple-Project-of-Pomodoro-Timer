package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/stop", handleStop)
	http.HandleFunc("/status", handleStatus)
	http.HandleFunc("/reset", handleReset)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Pomodoro Timer running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
