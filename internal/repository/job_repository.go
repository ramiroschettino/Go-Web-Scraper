package repository

import (
	"database/sql"
	"log"

	"github.com/ramiroschettino/Go-Web-Scraper/internal/models"
)

func SaveJobs(db *sql.DB, jobs []models.Job) {
	for _, job := range jobs {
		_, err := db.Exec(
			"INSERT INTO jobs (title, company, location, posted_date) VALUES ($1, $2, $3, $4)",
			job.Title, job.Company, job.Location, job.PostedDate,
		)
		if err != nil {
			log.Println("Error guardando trabajo:", err)
		}
	}
	log.Println("✅ Trabajos guardados en la base de datos")
}
