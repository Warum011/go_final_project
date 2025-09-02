package api

import (
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": "id not defined"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, task)
}
