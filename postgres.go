package main

import (
	"fmt"
	"os"
	"reflect"
)

type PostgresDSN struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

func NewPostgresDSN() PostgresDSN {
	dsn := PostgresDSN{
		User:     os.Getenv("POSTGRESQL_USER"),
		Password: os.Getenv("POSTGRESQL_PASSWORD"),
		Host:     os.Getenv("PERSON_SVC_GO_DB_SERVICE_HOST"),
		Port:     os.Getenv("POSTGRESQL_PORT"),
		Database: os.Getenv("POSTGRESQL_DATABASE"),
		SSLMode:  os.Getenv("POSTGRESQL_SSL"),
	}
	return dsn
}

func (dsn PostgresDSN) Valid() bool {
	v := reflect.ValueOf(dsn)
	for i := 0; i < v.NumField(); i++ {
		if len(v.Field(i).String()) == 0 {
			return false
		}
	}
	return true
}

func (dsn PostgresDSN) ToString() string {
	return fmt.Sprintf("user=%v password=%v host=%v port=%v database=%v sslmode=%v", dsn.User, dsn.Password, dsn.Host, dsn.Port, dsn.Database, dsn.SSLMode)
}
