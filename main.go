package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// heyo Returns a 200 with a greeting and its hostname header
func heyo(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost.localdomain"
	}
	w.Header().Add("x-hostname", hostname)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Heyo!\n")
}

func main() {
	listenPort := ":8080"
	lp := os.Getenv("LISTEN_PORT")
	if _, err := strconv.Atoi(lp); err == nil {
		listenPort = ":" + lp
	}
	http.HandleFunc("/", heyo)
	log.Printf("Listening on 0.0.0.0%v\n", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
