package main

import (
	"database/sql"
	"fmt"
	"resume-editor/go-backend/jobs"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//initialize database connection

	db, err := sql.Open("sqlite3", "./project.db")
	if err != nil {
		fmt.Println("Something wen wrong here")
	}
	defer db.Close()

	// database repository action

	fmt.Println("You have passed stage 1")

	// Every model will need to setup the DB

	jobs.Setup(db)

	jobs.CreateTable()

	/////////////////

	/////////////////

	/////////////////

	/////////////////

	router := gin.Default()

	router.GET("/jobs", jobs.GetAllJobs)
	router.POST("/jobs", jobs.AddJob)

	router.Run(":8084")

}
