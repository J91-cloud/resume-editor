package jobs

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func Setup(database *sql.DB) {
	db = database
}

func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS jobs(
	
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	date_applied TEXT NOT NULL,
	job_type TEXT NOT NULL
	
	
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table: ", err)

	}

}

type job struct {
	id           int
	name         string `json:name`
	date_applied string `json:date_applied`
	job_type     string `json:job_type`
}

func GetAllJobs(c *gin.Context) {
	var jobs []job

	query := "SELECT id,name,date_applied,job_type FROM jobs"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
	}
	for rows.Next() {
		var u job
		if err := rows.Scan(&u.id, &u.name, &u.job_type, *&u.date_applied); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
			return
		}
		jobs = append(jobs, u)
	}

}

func AddJob(c *gin.Context) {
	var newjob job
	if err := c.BindJSON(&newjob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO users (name, age, class) VALUES (?, ?, ?)", newjob.name, newjob.job_type, newjob.date_applied)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	newjob.id = int(id)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Resource created successfully.", "data": newjob})
}
