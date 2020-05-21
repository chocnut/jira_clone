package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chocnut/jira_clone/controllers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/ping", controllers.Ping)
	handler.HandleFunc("/projects", controllers.GetAllProject)

	err := http.ListenAndServe(port, handler)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
