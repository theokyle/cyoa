package main

import (
	"log"
	"net/http"
)

func (cfg *Config) handlerChapter(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("chapter")
	if path == "" || path == "/" {
		path = "intro"
	}

	if chapter, ok := cfg.story[path]; ok {
		err := cfg.template.Execute(w, chapter)
		if err != nil {
			log.Printf("error: %v", err)
			http.Error(w, "Unable to Render Template", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found", http.StatusBadRequest)
}
