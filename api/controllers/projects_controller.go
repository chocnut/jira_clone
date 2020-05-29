package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/chocnut/jira_clone/api/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {

	db, err := sqlx.Open("mysql", "root:password@tcp(127.0.0.1:3306)/jira_clone?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	projects := []models.Project{}

	db.Select(&projects, "SELECT * FROM projects")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
