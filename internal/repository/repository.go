package repository

import (
	"fmt"
	"log"

	"github.com/ramiroschettino/Go-Web-Scraper/internal/models"
)

func SaveScrapedData(scrapedData *models.ScrapedData) error {
	if DB == nil {
		log.Fatal("La base de datos no está inicializada.")
		return fmt.Errorf("base de datos no inicializada")
	}

	query := "INSERT INTO scraped_data (text, url) VALUES (?, ?)"
	_, err := DB.Exec(query, scrapedData.Text, scrapedData.URL)
	if err != nil {
		return err
	}

	return nil
}
