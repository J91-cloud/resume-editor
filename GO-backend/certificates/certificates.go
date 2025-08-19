package certificates

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

type certificates struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Date_applied string `json:"date_applied"`
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS certificates (
	ID PRIMARY KEY AUTO INCREMENT,
	Name TEXT NOT NULL,
	Date_applied TEXT NOT NULL
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating certificates tables", err)

	}

}

func GetAllCertificates(g *gin.Context) {
	var certificate []certificates

	query := `SELECT ID,Name,Date_applied from certificates`

	rows, err := db.Query(query)

	if err != nil {
		g.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false})
	}

	for rows.Next() {
		var u certificates
		if err := rows.Scan(&u.ID, &u.Name, &u.Date_applied); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error fetching the certificates database", "error": err.Error()})
			return
		}
		certificate = append(certificate, u)
	}
	g.JSON(http.StatusAccepted, gin.H{"success": true, "mesagge": "Successful fetch of the certificates database", "data": certificate})

}

func CreateCertificate(c *gin.Context) {
	var newCertificate certificates

	result, err := db.Exec("INSERT INTO certificates (name, date_applied) VALUES (?,?)", newCertificate.Name, newCertificate.Date_applied)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create certificate", "errpr": err.Error()})
	}

	id, err := result.LastInsertId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create certificate", "errpr": err.Error()})
	}

	newCertificate.ID = int(id)

	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Successfuly created new certificate", "data": newCertificate})
}
