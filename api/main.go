package main

import (
	"fmt"
	"os"

	"github.com/aweglteo/fullstack_development/api/external/grpcc"
	"github.com/aweglteo/fullstack_development/api/external/postgres"
	"github.com/aweglteo/fullstack_development/api/external/router"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	fmt.Printf("loading env ... ")
	loadEnv()
	fmt.Printf("DBHOST: %s\n", os.Getenv("DB_HOST"))

	r := router.NewRouter()
	// establish PostgreSQL connection
	// postgres.Connect()
	grpcc.Connect()

	defer grpcc.CloseConn()
	defer postgres.CloseConn()
	r.Run()
}
