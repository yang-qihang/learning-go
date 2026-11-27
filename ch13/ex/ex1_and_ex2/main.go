// How to verify the implementation?
//
// 1. Test the plain-text response:
// curl -i http://127.0.0.1:8080/time
//
// 2. Test the JSON response:
// curl -i -H "Accept: application/json" http://127.0.0.1:8080/time
package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)

	// Or more directly:
	// w.Write([]byte(currentTime))

	// the interface http.ResponseWriter also satisfies io.Writer

	fmt.Fprintln(w, currentTime)
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
