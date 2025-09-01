package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Warum011/go_final_project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		if err := writeJson(w, map[string]string{"error": "method not allowed"}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := writeJson(w, map[string]string{"error": "id not defined"}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
		return
	}

	nowStr := r.URL.Query().Get("now")
	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err := writeJson(w, map[string]string{"error": "invalid now format"}); err != nil {
				log.Printf("writeJson error: %v", err)
			}
			return
		}
	}

	task, err := db.GetTask(id)
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

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
				log.Printf("writeJson error: %v", err)
			}
			return
		}
	} else {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
				log.Printf("writeJson error: %v", err)

				return
			}
			task.Date = next
			if err := db.UpdateTask(task); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				if err := writeJson(w, map[string]string{"error": err.Error()}); err != nil {
					log.Printf("writeJson error: %v", err)
				}
				return
			}
		}
		if err := writeJson(w, map[string]string{}); err != nil {
			log.Printf("writeJson error: %v", err)
		}
	}
}
