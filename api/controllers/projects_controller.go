package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/chocnut/jira_clone/api/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllProjects(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/jira_clone?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT id, name, description, created_at, updated_at FROM projects")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		var project models.Project

		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.CreatedAt, &project.UpdatedAt); err != nil {
			log.Fatal(err)
		}

		projects = append(projects, project)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
