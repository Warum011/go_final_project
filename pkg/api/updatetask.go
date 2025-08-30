package api

import (
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := readJson(r.Body, &task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if task.ID == "" || task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": "not all fields are filled in"})
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, map[string]string{})
}
