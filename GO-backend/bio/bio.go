package bio

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bio struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Phone_number      string `json:"phone_number""`
	Address           string `json:"address"`
	Short_description string `json:"short_description"`
	Linkedln_link     string `json:"linkedln_link"`
}

var db *sql.DB

func Setup(database *sql.DB) {
	db = database
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS bio (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	phone_number TEXT NOT NULL,
	address TEXT NOT NULL,
	short_description TEXT NOT NULL,
	linkedln_link TEXT NOT NULL

	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating the bio table", err)
	}
}

func GetAllBIO(c *gin.Context) {
	var bio_profile []bio

	query := "SELECT * FROM bio;"

	result, err := db.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to get all of bio", "error": err.Error()})
	}

	for result.Next() {
		var u bio
		if err := result.Scan(&u.ID, &u.Name, &u.Phone_number, &u.Short_description, &u.Linkedln_link); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to get all of bio", "error": err.Error()})

		}

		bio_profile = append(bio_profile, u)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": "Successfuly got all of the bio", "data": bio_profile})
}

/*

Create a Update Function for this source

*/
