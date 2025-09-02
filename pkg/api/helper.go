package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Warum011/go_final_project/pkg/db"
)

func writeJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJson error: %v", err)
	}
}

func readJson(body io.ReadCloser, v interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(v)
}

func checkDate(task *db.Task) error {
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	_, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYYMMDD")
	}

	if task.Date < now.Format(dateFormat) {
		if task.Repeat != "" {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		} else {
			task.Date = now.Format(dateFormat)
		}
	}
	return nil
}
