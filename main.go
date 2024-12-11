package main

import (
	"fmt"
	"log"
	"net/http"
)
import (
	"alexdev2001/zoo_database_crud_api/database_connect"
	"alexdev2001/zoo_database_crud_api/routes"
)

func main() {
	fmt.Println("Hello, programmer. Welcome to the zoo API")

	// set up the dabase connection
	db, err := database_connect.DBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// set up routes
	initializedRoutes := routes.Routes(db)
	fmt.Println("Listening on port 8080")
	err = http.ListenAndServe(":8080", initializedRoutes)
	if err != nil {
		log.Fatal(err)
	}
}
