package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	filename := flag.String("file", "gopher.json", "A json file with story content")
	port := flag.Int("port", 8080, "Port to run application")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error opening file: %s", *filename)
	}
	defer file.Close()

	story, err := JsonStory(file)
	if err != nil {
		fmt.Printf("error decoding story: %s", err)
	}

	tmpl := template.Must(template.ParseFiles("template.html"))

	cfg := Config{
		story:    story,
		template: tmpl,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", cfg.handlerChapter)
	mux.HandleFunc("/{chapter}", cfg.handlerChapter)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

	fmt.Printf("Starting application on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

type Config struct {
	story    Story
	template *template.Template
}
