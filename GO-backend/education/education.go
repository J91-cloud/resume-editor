package education

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
	query := `CREATE TABLE IF NOT EXISTS education (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	date_started TEXT NOT NULL,
	date_ended TEXT NOT NULL,
	services TEXT NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table: ", err)

	}

}

type education struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Date_started string `json:"date_started"`
	Date_ended   string `json:"date_ended"`
	// Services     []string `json:"services`
}

func GetAllEducation(c *gin.Context) {
	var edu []education

	query := "SELECT * FROM education"

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
	}

	for rows.Next() {
		var u education
		if err := rows.Scan(&u.ID, &u.Name, &u.Date_started, &u.Date_ended); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
			return
		}

		edu = append(edu, u)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "Successfully fetched all education", "data": edu})

}
