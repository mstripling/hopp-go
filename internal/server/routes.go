package server

import (
	"encoding/json"
	"log"
	"net/http"
  "fmt"
  "hopp/internal/util"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
  mux.HandleFunc("/json", s.HashHandler)

	mux.HandleFunc("/health", s.healthHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) HashHandler(w http.ResponseWriter, r *http.Request) {
 // Only accept POST requests
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Decode the JSON request body into the struct
    var payload util.Payload

    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        http.Error(w, fmt.Sprintf("Invalid JSON: %s", err), http.StatusBadRequest)
        return 
    }
    if payload.Plain == nil {
      payload.Plain = make(map[string]interface{})
    }
    
    if payload.Hash == nil {
      payload.Hash = make(map[string]interface{})
    }

  // Now call the HashFunction to process the data
	processedPayload, err := util.Hash(payload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing data: %s", err), http.StatusInternalServerError)
		return
	}

	// Return the merged data as a single JSON object
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(processedPayload); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %s", err), http.StatusInternalServerError)
		return
	}
}

