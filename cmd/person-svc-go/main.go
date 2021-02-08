package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/laetho/person-svc-go/internal/pkg/handlers"
	"github.com/laetho/person-svc-go/internal/pkg/postgres"
	"github.com/laetho/person-svc-go/pkg/persons"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Concurrency safe Postgres connection pool. Going global here for such
// a simple application.
var DBCP *pgxpool.Pool
var DSN postgres.PostgresDSN

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := godotenv.Load()
	if err != nil {
		log.Info("Unable to load .env file, ignoring...")
	}

	DSN = postgres.NewPostgresDSN()
	if !DSN.Valid() {
		log.Fatalf("Unable to Generate Postgresql DSN. Missing environment variables?")
	}
}

func main() {
	var err error

	DBCP, err = pgxpool.Connect(context.Background(), DSN.ToString())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)

	}
	defer DBCP.Close()

	// Run migration
	conn, err := DBCP.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Unable to get DB connection: %v", err)
	}
	defer conn.Release()

	_, err = conn.Query(context.Background(), persons.PersonMigration)
	if err != nil {
		panic(err)
	}

	port := 8080
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/status", handlers.StatusHandler)
	http.HandleFunc("/persons", handlers.GetPersonsHandler(DBCP))
	http.HandleFunc("/create", handlers.CreatePersonHandler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
