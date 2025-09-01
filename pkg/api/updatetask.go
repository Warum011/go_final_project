package api

import (
	"log"
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := readJson(r.Body, &task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	if task.ID == "" || task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": "not all fields are filled in"}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}
	if err := writeJson(w, map[string]string{}); err != nil {
		log.Printf("writeJson error: %v", err)
	}
}
