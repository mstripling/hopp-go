package server

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "io"
    "hopp/internal/util"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)
  mux.HandleFunc("/ping", s.VendorPingHandler)
  mux.HandleFunc("/buyer", s.BuyerBidHandlerTest)
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

func (s *Server) BuyerBidHandlerTest (w http.ResponseWriter, r *http.Request) {
  resp := make(map[string]interface{})
  resp["pingID"] = 1234567890
  resp["bid"] = 5

  jsonResp, err := json.Marshal(resp)
  if err != nil {
    log.Fatalf("Error handling JSON marshal. Err: %v", err)
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

    err = util.Initialize(w, r, vendorPayload)
    if err != nil {
      http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
      return 
  }
  
  // Now call the TransformAndFormat func to process the data
	processedPayload, err := util.TransformAndFormat(vendorPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing data: %s", err), http.StatusInternalServerError)
		return
	}

  // Return positive response to vendor
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Request processed successfully.....\nStandby for Buyer Response\n"))
  fmt.Println(processedPayload)

  resp, err := util.Ping(r, processedPayload)
  if err != nil {
    http.Error(w, fmt.Sprintf("HTTP request error: %s", err), http.StatusBadRequest)
    return
  }

  w.WriteHeader(resp.StatusCode)
  _, err = io.Copy(w, resp.Body)
  if err != nil {
    http.Error(w, fmt.Sprintf("Error copying response body: %s", err), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()
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
