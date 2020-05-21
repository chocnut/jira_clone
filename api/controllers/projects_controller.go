package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chocnut/jira_clone/api/models"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	var Projects = []models.Project{
		models.Project{ID: 1, Name: "Sample Project 1", CreatedAt: time.Now()},
		models.Project{ID: 2, Name: "Sample Project 2", CreatedAt: time.Now()},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Projects)
}
