package skills

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

type skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS skills (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
	
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating table", err)
	}
}

func GetAllSkills(c *gin.Context) {
	var skills []skill
	query := `SELECT id, name FROM skills`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Sommething went wrong fetchting skills database", err)
	}

	for rows.Next() {
		var u skill
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error fetching jobs ", "error": err.Error()})
		}
		skills = append(skills, u)
	}

	c.JSON(http.StatusInternalServerError, gin.H{"success": "All good from here", "message": "Fetched jobs correctly", "data": skills})

}

func CreateSkill(c *gin.Context) {
	var newSkill skill

	result, err := db.Exec("INSERT INTO skills (name) VALUES (?)", newSkill.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating new skill", "error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	newSkill.ID = int(id)

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "successfully created a new skill", "data": newSkill})

}

func DeleteSkill(c *gin.Context) {
	var skillid = c.Param("id")

	idString, err := strconv.Atoi(skillid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting skill.", "error": err.Error()})
		return
	}

	result, err := db.Exec("DELETE FROM skills where id = ?", idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting skill.", "error": err.Error()})
		return
	}
	rowAffected, err := result.RowsAffected()

	if err != nil || rowAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting skill.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"success": true, "message": "Successfuly deleted skill"})

}
