package main

import (
	"log"
	"net/http"

	"github.com/chocnut/jira_clone/controllers"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/ping", controllers.Ping)
	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
