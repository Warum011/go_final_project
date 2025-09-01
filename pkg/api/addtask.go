package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

func addTaskHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var task db.Task

	if err := readJson(r.Body, &task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": "title not defined"}); err != nil {
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

	id, err := db.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	if err := writeJson(w, map[string]string{"id": fmt.Sprint(id)}); err != nil {
		log.Printf("failed to write JSON response: %v", err)
	}
}
