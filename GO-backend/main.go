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

	jobs.Setup(db)

	fmt.Println("You have passed stage 2")

	jobs.CreateTable()

	fmt.Println("You have passed stage 3")

	r := gin.Default()

	r.GET("/jobs", jobs.GetAllJobs)
	r.POST("/jobs", jobs.AddJob)

}
