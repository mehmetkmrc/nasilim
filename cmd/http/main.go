package main

import (
	"log"

	"github.com/mehmetkmrc/nasilim.git/internal/adapter/secondary/storage/postgres"
)

func main() {
	
	connString := "postgres://username:password@localhost:5432/dbname?sslmode=disable"
	postgres.NewPostgresUserRepository(connString)
	log.Fatal()

}