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
  mux.HandleFunc("/json", s.VendorPingHandler)

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

func (s *Server) VendorPingHandler(w http.ResponseWriter, r *http.Request) {
    // Decode the JSON request body into the struct
    var vendorPayload util.RawPayload

    err := json.NewDecoder(r.Body).Decode(&vendorPayload)
    if err != nil {
        http.Error(w, fmt.Sprintf("Invalid JSON: %s", err), http.StatusBadRequest)
        return 
    }

    if vendorPayload.Endpoint == ""{
    http.Error(w, fmt.Sprintf("Invalid endpoint: %s", vendorPayload.Endpoint), http.StatusBadRequest)
    }
    if vendorPayload.Plain == nil {
      vendorPayload.Plain = make(map[string]interface{})
    }
    
    if vendorPayload.Hash == nil {
      vendorPayload.Hash = make(map[string]interface{})
    }

  // Now call the TransformAndFormat func to process the data
	processedPayload, err := util.TransformAndFormat(vendorPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing data: %s", err), http.StatusInternalServerError)
		return
	}

  // Return positive response to vendor
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Request processed successfully.....\nStandby for Buyer Response\n:v"))
  fmt.Println(processedPayload)


/*
	// Return the merged data as a single JSON object
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(processedPayload)
  if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %s", err), http.StatusInternalServerError)
		return
	}
*/
}

