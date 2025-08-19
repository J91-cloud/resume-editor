package main

import (
	"database/sql"
	"fmt"
	"resume-editor/go-backend/bio"
	"resume-editor/go-backend/certificates"
	"resume-editor/go-backend/education"
	"resume-editor/go-backend/jobs"
	"resume-editor/go-backend/projects"
	"resume-editor/go-backend/skills"

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

	router := gin.Default()

	{

		jobs.Setup(db)

		jobs.CreateTable()

		job_endpoints := router.Group("/jobs")

		job_endpoints.GET("", jobs.GetAllJobs)
		job_endpoints.POST("", jobs.AddJob)
		job_endpoints.DELETE("/:id", jobs.DeleteJob)

	}

	// skills group of routes

	{

		skills.Setup(db)
		skills.CreateTable()

		skill_endpoint := router.Group("/skills")
		skill_endpoint.GET("", skills.GetAllSkills)
		skill_endpoint.POST("", skills.CreateSkill)
		skill_endpoint.DELETE("", skills.DeleteSkill)
	}

	// projects group of routes

	{
		projects.Setup(db)
		projects.CreateTable()

		project_endpoint := router.Group("/projects")
		project_endpoint.GET("", projects.GetAllProjects)
	}

	//certificate endpoints

	{
		certificates.Setup(db)
		certificates.CreateTable()

		certificate_endpoints := router.Group("/certificates")
		certificate_endpoints.GET("", certificates.GetAllCertificates)
		certificate_endpoints.POST("", certificates.CreateCertificate)
		certificate_endpoints.DELETE("/:id", certificates.DeleteCertificate)
	}

	// bio groupe of routes

	{
		bio.Setup(db)
		bio.CreateTable()

		bio_endpoints := router.Group("/bio")
		bio_endpoints.GET("", bio.GetAllBIO)
	}

	{
		education.Setup(db)
		education.CreateTable()

		education_endpoints := router.Group("/education")
		education_endpoints.GET("", education.GetAllEducation)
	}

	router.Run(":8084")

}
