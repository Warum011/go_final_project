package api

import (
	"net/http"
	"time"

	"github.com/Warum011/go_final_project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJson(w, map[string]string{"error": "method not allowed"})
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": "id not defined"})
		return
	}

	nowStr := r.URL.Query().Get("now")
	var now time.Time
	if nowStr == "" {
		now = time.Now()
	} else {
		tm, err := time.Parse(dateFormat, nowStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJson(w, map[string]string{"error": "invalid now format"})
			return
		}
		now = tm
	}

	task, err := db.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			writeJson(w, map[string]string{"error": err.Error()})
			return
		}
		writeJson(w, map[string]string{})
		return
	}

	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}
	task.Date = next

	if err := db.UpdateTask(task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, map[string]string{})
}
