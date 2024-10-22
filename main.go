package main

import (
	"fmt"

	ascii "ascii_web/Objects"
)

func main() {
	fmt.Println("http://localhost:8080/")
	var Server ascii.Server

	Server.Run()
}
