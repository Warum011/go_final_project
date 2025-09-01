package api

import (
	"log"
	"net/http"

	"github.com/Warum011/go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": "method not allowed"}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	tasks, err := db.Tasks(50) // в параметре максимальное количество записей
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}
	if err := writeJson(w, TasksResp{Tasks: tasks}); err != nil {
		log.Printf("writeJson error: %v", err)
	}
}
