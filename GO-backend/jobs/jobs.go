package jobs

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Date_applied string `json:"date_applied"`
	Job_type     string `json:"job_type"`
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

	result, err := db.Exec("INSERT INTO jobs (name, date_applied, job_type) VALUES (?, ?, ?)", newjob.Name, newjob.Date_applied, newjob.Job_type)
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

func DeleteJob(c *gin.Context) {
	var deletedJobId = c.Param("id")

	idString, err := strconv.Atoi(deletedJobId)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	result, err := db.Exec("DELETE from job where id = ? ", idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "did not delete job successfuly", "error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		// 3. Check if any rows were actually deleted
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Job not found or already deleted."})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "Successfuly delete your job"})

}
