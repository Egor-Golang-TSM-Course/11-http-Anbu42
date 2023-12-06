package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("Received request to /")
	w.Write([]byte("Hello, welcome to the server!"))

}

func showTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/time" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("Received request to /time")
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	w.Write([]byte(fmt.Sprintf("Current time: %s", currentTime)))
}

func startServer() {
	http.HandleFunc("/", root)
	http.HandleFunc("/time", showTime)
	log.Println("Server is listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func main() {
	startServer()
}
