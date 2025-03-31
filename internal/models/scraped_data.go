package models

type ScrapedData struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	URL  string `json:"url"`
}
