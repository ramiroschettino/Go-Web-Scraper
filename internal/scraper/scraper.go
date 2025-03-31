package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/db"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/models"
)

func ScrapeWebsite(url string) ([]map[string]string, error) {
	c := colly.NewCollector()
	var results []map[string]string

	c.OnHTML("a", func(e *colly.HTMLElement) {
		text := e.Text
		link := e.Attr("href")

		data := models.ScrapedData{
			URL:  link,
			Text: text,
		}

		err := db.SaveScrapedData(data)
		if err == nil {
			results = append(results, map[string]string{"text": text, "url": link})
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("error al visitar la página: %v", err)
	}

	log.Println("✅ Scraping finalizado")
	return results, nil
}
