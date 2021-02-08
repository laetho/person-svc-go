package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/laetho/person-svc-go/pkg/persons"
	"log"
	"net/http"
	"os"
)

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = fmt.Fprintf(w, "{\"status\": \"UP\"}")
}

func GetPersonsHandler(pool *pgxpool.Pool) http.HandlerFunc {
	fn := func (w http.ResponseWriter, req *http.Request) {
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to get DB connection: %v\n", err)
		}
		defer conn.Release()

		var persons persons.Persons
		err = pgxscan.Select(context.Background(), conn, &persons, "select * from person")
		if err != nil {
			log.Printf("Query failed: %v", err)
		}
		out, _ := json.Marshal(persons)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		_, _ = fmt.Fprintf(w, string(out))
	}
	return http.HandlerFunc(fn)
}


func CreatePersonHandler(w http.ResponseWriter, req *http.Request) {

}
