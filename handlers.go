package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/georgysavva/scany/pgxscan"
)

func statusHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = fmt.Fprintf(w, "{\"status\": \"UP\"}")
}

func getPersonsHandler(w http.ResponseWriter, req *http.Request) {
	conn , err := dbcp.Acquire(context.Background())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to get DB connection: %v\n", err)
	}
	var persons []*Person
	err = pgxscan.Select(context.Background(), conn, &persons,"select * from person" )
	if err != nil {
		log.Printf("Query failed: %v", err)
	}
	out, _ := json.Marshal(persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = fmt.Fprintf(w, string(out))
}

func createPersonHandler(w http.ResponseWriter, req *http.Request) {

}