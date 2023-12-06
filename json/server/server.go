package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	fmt.Println("Server is listening on :8080")
	err := http.ListenAndServe(":8080", routes())
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func main() {
	startServer()
}
