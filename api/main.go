package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"rockonsoft.com/gx-state-api/api"
	"rockonsoft.com/gx-state-api/db"
)

func main() {
	log.Print("server has started")
	//start the db
	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error: %v", err)
		panic("error starting the database")

	}
	//get the router of the API by passing the db
	router := api.StartAPI(pgdb)
	//get the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(fmt.Sprintf("Server listening on: %v", port))
	//pass the router and start listening with the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
		return
	} else {
		log.Printf("Server still listening on: %v", port)
	}
}
