package main

import (
	"hexagonal-architecture-golang/src/bootstrap/server"
)

func main() {
	// database, err := database.NewDatabaseConnection()
	// if err != nil {
	// 	panic(err)
	// }
	// database.Initialice()
	server, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	server.Initialice()

}
