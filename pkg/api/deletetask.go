package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": "id not defined"}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		if errors.Is(err, db.ErrTaskNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}
	if err := writeJson(w, map[string]string{}); err != nil {
		log.Printf("writeJson error: %v", err)
	}
}
