package projects

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

type project struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Short_description string `json:"short_description"`
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS projects (
	ID PRIMARY KEY AUTO INCREMENT,
	Name TEXT NOT NULL,
	Short_description TEXT NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating projects database", err)

	}
}

func GetAllProjects(c *gin.Context) {
	var projects []project
	query := `SELECT ID,Name,Short_description from projects`
	rows, err := db.Query(query)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)

	}

	for rows.Next() {
		var u project
		if err := rows.Scan(&u.ID, &u.Name, &u.Short_description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": "Error getting from the projects table", "message": "fetched projects table", "error": err.Error()})

		}

		projects = append(projects, u)
	}

	c.JSON(http.StatusInternalServerError, gin.H{"success": "Successful get from the projects table", "message": "fetched projects table", "data": projects})

}

func CreateProject(c *gin.Context) {
	var newProject project

	result, err := db.Exec("INSERT INTO projects (name, short_description) VALUES (?,?)", newProject.Name, newProject.Short_description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create new project", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create new project", "error": err.Error()})
		return
	}

	newProject.ID = int(id)

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "successfuly created a new project", "data": newProject})

}

func DeleteProject(c *gin.Context) {
	var projectId = c.Param("id")

	idString, err := strconv.Atoi(projectId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting project.", "error": err.Error()})
		return
	}

	result, err := db.Exec("DELETE FROM projects where id = ?", idString)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting skill.", "error": err.Error()})
		return
	}

	rowAffected, err := result.RowsAffected()
	if err != nil || rowAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting skill.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "Successfuly deleted project"})

}
