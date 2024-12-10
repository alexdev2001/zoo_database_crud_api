package main

import (
	"fmt"
)
import (
	"alexdev2001/zoo_database_crud_api/database_connect"
)

func main() {
	fmt.Println("Hello, programmer. Welcome to the zoo API")

	// set up the dabase connection
	database_connect.DBConnection()
	
}
