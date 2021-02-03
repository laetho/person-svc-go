package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"status\": \"UP\"}")
}

func main() {
	port := 8080
	http.HandleFunc("/status", statusHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
