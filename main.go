package main

import (
	"fmt"
	"log"

	"github.com/mehmetkmrc/nasilim.git/server"
)

func main() {
	fmt.Println("Go runs here!")
	if _, err := server.NewServer("0.0.0.0:8087"); err != nil{
		log.Fatal("Failed to start server")
	}
}