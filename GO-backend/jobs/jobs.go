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
    Id           int    `json:"id"`           // Add quotes around "id"
    Name         string `json:"name"`         // Add quotes around "name"  
    Date_applied string `json:"date_applied"` // Add quotes around "date_applied"
    Job_type     string `json:"job_type"`     // Add quotes around "job_type"
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
if err := rows.Scan(&u.Id, &u.Name, &u.Date_applied, &u.Job_type); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
return
 }
jobs = append(jobs, u)
 }
c.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "Jobs retrieved successfully.", "data": jobs})
}

func AddJob(c *gin.Context) {
	var newjob job
	if err := c.BindJSON(&newjob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO jobs (name, date_applied, job_type) VALUES (?, ?, ?)", newjob.Name,newjob.Date_applied, newjob.Job_type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	newjob.Id = int(id)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Resource created successfully.", "data": newjob})
}
