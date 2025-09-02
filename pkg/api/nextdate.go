package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

var (
	ErrInvalidRepeat     = errors.New("rules are not followed or skipped")
	ErrInvalidDateFormat = errors.New("invalid date format, expected YYYYMMDD")
	ErrUsupportedFormat  = errors.New("format is not supported")
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {

	startDate, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", ErrInvalidDateFormat
	}

	if repeat == "" {
		return "", ErrInvalidRepeat
	}

	steps := strings.Split(repeat, " ")

	switch steps[0] {
	case "y":
		if len(steps) > 1 {
			return "", ErrInvalidRepeat
		}

		for {
			startDate = startDate.AddDate(1, 0, 0)
			if startDate.After(now) {
				break
			}
		}
		return startDate.Format(dateFormat), nil

	case "d":
		if len(steps) == 1 {
			return "", ErrInvalidRepeat
		} else {
			days, err := strconv.Atoi(steps[1])
			if err != nil {
				return "", ErrInvalidRepeat
			}

			if days <= 0 || days > 400 {
				return "", ErrInvalidRepeat
			}

			for {
				startDate = startDate.AddDate(0, 0, days)
				if startDate.After(now) {
					break
				}
			}
			return startDate.Format(dateFormat), nil
		}

	case "w", "m":
		return "", ErrUsupportedFormat

	default:
		return "", ErrInvalidDateFormat
	}
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeJson(w, map[string]string{"error": "method not allowed"})
		return
	}

	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	if dateStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": "date missing"})
		return
	}

	repeatStr := r.FormValue("repeat")
	if repeatStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": "repeat missing"})
	}

	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			writeJson(w, map[string]string{"error": ErrInvalidDateFormat.Error()})
			return
		}
	}

	next, err := NextDate(now, dateStr, repeatStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte(next)); err != nil {
		log.Printf("failed to write response in /api/nextdate: %v", err)
	}
}
