package api

import (
	"fmt"
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
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if task.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": "title not defined"})
		return
	}

	if err := checkDate(&task); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, map[string]string{"id": fmt.Sprint(id)})
}
