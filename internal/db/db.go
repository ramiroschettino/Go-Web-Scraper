package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/models"
)

var DB *sql.DB

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Error conectando a la base de datos: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("❌ No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("✅ Conectado a PostgreSQL exitosamente")
}

// SaveScrapedData guarda los datos scrapeados en la base de datos
func SaveScrapedData(data models.ScrapedData) error {
	query := `INSERT INTO scraped_data (url, content) VALUES ($1, $2)`
	_, err := DB.Exec(query, data.URL, data.Content)
	if err != nil {
		log.Printf("❌ Error guardando en la base de datos: %v", err)
		return err
	}

	log.Println("✅ Datos guardados en la base de datos")
	return nil
}
