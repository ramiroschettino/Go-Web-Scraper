package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/models"
	"github.com/ramiroschettino/Go-Web-Scraper/internal/repository"
)

func ScrapeWebsite(url string) ([]map[string]string, error) {
	c := colly.NewCollector()
	var results []map[string]string

	c.OnRequest(func(r *colly.Request) {
		log.Printf("🌍 Visitando: %s", r.URL.String())
	})

	// Scrapeando enlaces (a)
	c.OnHTML("a", func(e *colly.HTMLElement) {
		text := e.Text
		link := e.Attr("href")

		link = e.Request.AbsoluteURL(link)
		if link == "" {
			return
		}

		log.Printf("Enlace encontrado: %s - %s", text, link)

		data := models.ScrapedData{
			URL:  link,
			Text: text,
		}

		// Pasamos el puntero a la función SaveScrapedData
		err := repository.SaveScrapedData(&data)
		if err == nil {
			results = append(results, map[string]string{"text": text, "url": link})
		}
	})

	// Scrapeando encabezados (h1, h2, h3, etc.)
	c.OnHTML("h1, h2, h3, h4, h5, h6", func(e *colly.HTMLElement) {
		text := e.Text
		log.Printf("Encabezado encontrado: %s", text)

		results = append(results, map[string]string{"type": e.Name, "text": text})
	})

	// Scrapeando párrafos (p)
	c.OnHTML("p", func(e *colly.HTMLElement) {
		text := e.Text
		log.Printf("Párrafo encontrado: %s", text)

		results = append(results, map[string]string{"type": "p", "text": text})
	})

	// Scrapeando imágenes (img)
	c.OnHTML("img", func(e *colly.HTMLElement) {
		imgSrc := e.Attr("src")
		altText := e.Attr("alt")

		imgSrc = e.Request.AbsoluteURL(imgSrc)
		if imgSrc == "" {
			return
		}

		log.Printf("Imagen encontrada: %s - %s", imgSrc, altText)

		results = append(results, map[string]string{"type": "img", "src": imgSrc, "alt": altText})
	})

	// Scrapeando listas (ul, ol)
	c.OnHTML("ul, ol", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, li *colly.HTMLElement) {
			listItem := li.Text
			log.Printf("Elemento de lista encontrado: %s", listItem)

			results = append(results, map[string]string{"type": "list_item", "text": listItem})
		})
	})

	// Scrapeando divs y spans
	c.OnHTML("div, span", func(e *colly.HTMLElement) {
		text := e.Text
		log.Printf("Texto encontrado en %s: %s", e.Name, text)

		results = append(results, map[string]string{"type": e.Name, "text": text})
	})

	// Scrapeando tablas (table)
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, tr *colly.HTMLElement) {
			var row []string
			tr.ForEach("td", func(i int, td *colly.HTMLElement) {
				row = append(row, td.Text)
			})
			if len(row) > 0 {
				log.Printf("Fila de tabla encontrada: %v", row)
				results = append(results, map[string]string{"type": "table_row", "content": fmt.Sprintf("%v", row)})
			}
		})
	})

	// Scrapeando metatags (meta)
	c.OnHTML("meta", func(e *colly.HTMLElement) {
		name := e.Attr("name")
		property := e.Attr("property")
		content := e.Attr("content")
		if name != "" {
			log.Printf("Meta tag encontrado (name): %s = %s", name, content)
			results = append(results, map[string]string{"type": "meta", "name": name, "content": content})
		}
		if property != "" {
			log.Printf("Meta tag encontrado (property): %s = %s", property, content)
			results = append(results, map[string]string{"type": "meta", "property": property, "content": content})
		}
	})

	// Scrapeando formularios (form)
	c.OnHTML("form", func(e *colly.HTMLElement) {
		e.ForEach("input, select, textarea", func(i int, elem *colly.HTMLElement) {
			name := elem.Attr("name")
			value := elem.Attr("value")
			if name != "" {
				log.Printf("Formulario encontrado: %s = %s", name, value)
				results = append(results, map[string]string{"type": "form_element", "name": name, "value": value})
			}
		})
	})

	// Scrapeando contenido de script o style
	c.OnHTML("script, style", func(e *colly.HTMLElement) {
		content := e.Text
		log.Printf("Contenido de %s encontrado: %s", e.Name, content[:50])
		results = append(results, map[string]string{"type": e.Name, "content": content[:50]})
	})

	// Scrapeando comentarios
	c.OnHTML("<!--", func(e *colly.HTMLElement) {
		comment := e.Text
		log.Printf("Comentario encontrado: %s", comment)
		results = append(results, map[string]string{"type": "comment", "content": comment})
	})

	// Scrapeando cualquier otro texto
	c.OnHTML("*", func(e *colly.HTMLElement) {
		text := e.Text
		if text != "" {
			log.Printf("Texto encontrado en %s: %s", e.Name, text)
			results = append(results, map[string]string{"type": e.Name, "text": text})
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("error al visitar la página: %v", err)
	}

	log.Println("✅ Scraping finalizado")
	return results, nil
}
