package main

import (
	"fmt"
	"io"
	"net/http"
)

func sendRequest(url, route string) {
	fmt.Printf("Sending request to %s\n", url)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("Response from %s:\n%s\n\n", route, body)
}

func main() {
	routes := []string{"/", "/time", "/nonexistent"}

	serverAddress := "http://localhost:8080"

	for _, route := range routes {
		sendRequest(serverAddress+route, route)
	}
}
