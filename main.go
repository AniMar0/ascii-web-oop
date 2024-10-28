package main

import (
	"fmt" 

	ascii "ascii_web/Objects" 
)

func main() {
	// Print the URL where the server will be accessible
	fmt.Println("http://localhost:8080/")

	//Server is a struct that contains server methods
	var Server ascii.Server

	// Call the Run method to start the server
	Server.Run()
}
