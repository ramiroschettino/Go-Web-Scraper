package service

import (
	"github.com/ramiroschettino/Go-Web-Scraper/internal/scraper"
)

func Scrape(url string) ([]map[string]string, error) {
	return scraper.ScrapeWebsite(url)
}
