package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ramiroschettino/Go-Web-Scraper/internal/scraper"
)

type ScrapeRequest struct {
	URL string `json:"url"`
}

func ScrapeHandler(w http.ResponseWriter, r *http.Request) {
	var req ScrapeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "❌ Error al leer el JSON", http.StatusBadRequest)
		return
	}

	log.Printf("🔎 Scrapeando: %s", req.URL)

	results, err := scraper.ScrapeWebsite(req.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("❌ Error en el scraping: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Println("🔍 Resultados del scraping:", results)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
