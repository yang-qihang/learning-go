// How to verify the implementation?
//
// 1. Test the plain-text response:
// curl -i http://127.0.0.1:8080/time
//
// 2. Test the JSON response:
// curl -i -H "Accept: application/json" http://127.0.0.1:8080/time
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	// currentTime := time.Now().Format(time.RFC3339)
	accept := r.Header.Get("Accept")
	now := time.Now()

	type CustomTimeFormat struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}

	if accept == "application/json" {
		nowCustom := CustomTimeFormat{
			DayOfWeek:  now.Weekday().String(),
			DayOfMonth: now.Day(),
			Month:      now.Month().String(),
			Year:       now.Year(),
			Hour:       now.Hour(),
			Minute:     now.Minute(),
			Second:     now.Second(),
		}

		// pay attention to the relative order between:
		//  1. setting the header
		//  2. encoding the information (i.e. setting it)
		//
		// BUT:
		//  You must set the response headers before writing the response body,
		//  not necessarily before creating the encoder.
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)
		err := encoder.Encode(nowCustom)
		// if err != nil {
		// 	panic(err)
		// }
		// Prefer meaningful handling of errors
		if err != nil {
			slog.Error("Failed to encode response", "error", err)
			return
		}
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, now.Format(time.RFC3339))
	}
}

func RequestLogger(next http.Handler) http.Handler {
	logger := slog.New(
		slog.NewJSONHandler(os.Stderr, nil),
	)

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)

			if err != nil {
				logger.Error("Failed to extract IP", "error", err)
			} else {
				logger.Info("Incoming request", "ip", ip)
			}

			next.ServeHTTP(w, r)
		},
	)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /time", getTime)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		// For exercise 1, it's directly mux
		Handler: RequestLogger(mux),
	}

	log.Println("Server is starting on port 8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
