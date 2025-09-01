package api

import (
	"log"
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": "id not defined"}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	if err := writeJson(w, task); err != nil {
		log.Printf("writeJson error: %v", err)
	}
}
