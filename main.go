package main

import (
	"log"

	"github.com/ramiroschettino/Go-Web-Scraper/cmd"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/db"
)

func main() {
	log.Println("🚀 Iniciando Go Web Scraper...")
	db.Connect()
	cmd.StartServer()
}
