package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chocnut/jira_clone/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Fatal("$PORT must be set")
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/ping", controllers.Ping)
	err := http.ListenAndServe(port, handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
