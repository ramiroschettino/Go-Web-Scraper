package db

import (
	"log"

	"github.com/ramiroschettino/Go-Web-Scraper/internal/models"
)

func SaveScrapedData(data models.ScrapedData) error {
	query := `INSERT INTO scraped_data (text, url) VALUES ($1, $2) ON CONFLICT (url) DO NOTHING`
	_, err := DB.Exec(query, data.Text, data.URL)
	if err != nil {
		log.Printf("❌ Error al insertar en la base de datos: %v", err)
		return err
	}

	log.Printf("✅ Guardado en la DB: %s", data.URL)
	return nil
}
