package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
  "log"

	_ "github.com/joho/godotenv/autoload"

	"hopp/internal/database"
)

func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()

        // Create a custom response writer to capture status code
        rw := &responseWriter{w, http.StatusOK}

        // Call the next handler in the chain
        next.ServeHTTP(rw, r)

        // Log the response details
        log.Printf("[%s] %s %s - Status: %d - Duration: %v", time.Now().Format(time.RFC3339), r.Method, r.URL.Path, rw.statusCode, time.Since(startTime))
    })
}

// responseWriter is a custom wrapper around http.ResponseWriter to capture status codes
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
    rw.statusCode = statusCode
    rw.ResponseWriter.WriteHeader(statusCode)
}


type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
 
	NewServer := &Server{
		port: port,
		db: database.New(),
	}

  log.Printf("Starting server on port %d...", NewServer.port)

  // Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      logMiddleware(NewServer.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}