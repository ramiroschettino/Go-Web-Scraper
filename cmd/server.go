package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/handler"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/scrape", handler.ScrapeHandler).Methods("POST")

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
