package server

import (
    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "io"
    "hopp-go/internal/util"
    "time"

    "hopp-go/internal/database"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HomeHandler)		// web interface testing
	mux.HandleFunc("/ping", s.VendorPingHandler)	// endpoint vendors use
	mux.HandleFunc("/buyer", s.BuyerBidHandlerTest) // handles web interface requests
	mux.HandleFunc("/health", s.healthHandler) 	// Provides Database connection stats
	mux.HandleFunc("/hello", s.HelloWorldHandler) 	// returns hellow world as json
	mux.HandleFunc("/login", s.LoginHandler) 	// Unused currently
	mux.HandleFunc("/register", s.RegisterHandler) 	// Unused currently
	mux.HandleFunc("/logout", s.LogoutHandler) 	// Unused currently
	mux.HandleFunc("/protected", s.ProtectedHandler) 	// Unused currently

	return mux
}

func (s *Server) HomeHandler (w http.ResponseWriter, r *http.Request) {
	path := "frontend/public/index.html"
	http.ServeFile(w, r, path)
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
	//for local testing
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON: %s", err), http.StatusBadRequest)
        	return 
    	}
	resp := make(map[string]interface{})
	if gender, ok := data["gender"]; ok && gender.(string) == "Female" {
		resp["bid"] = 5
	} else{resp["bid"] = 4}
	resp["pingID"] = 1234567890


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
    
	// Reference by pointer to init payload
	err = util.Initialize(w, r, &vendorPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
	return 
    	}
  
	// Create new payload that has been processed
	processedPayload, err := util.TransformAndFormat(&vendorPayload)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing data: %s", err), http.StatusInternalServerError)
	return
	}	

	if vendorPayload.Test == true {
		prettyjsonData, err := json.MarshalIndent(processedPayload, "","  ")
		if err != nil {
        	http.Error(w, fmt.Sprintf("Error marshalling data: %v", err), http.StatusInternalServerError)
        	return
		}
		w.Write([]byte("Request processed successfully.....\n"))
		w.Write([]byte("Your Hopp payload (prettified):\n"))
		w.Write([]byte(prettyjsonData))
		w.Write([]byte("\nStandby for Buyer Response\n"))
    	}
   
	resp, err := util.Ping(r, &processedPayload, vendorPayload.Endpoint)
	if err != nil {
		http.Error(w, fmt.Sprintf("HTTP request error: %s", err), resp.StatusCode)
		return
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error copying response body: %s", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
}


func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
     if r.Method == http.MethodGet {
        // Serve the HTML form
        path := "frontend/public/register.html"
        http.ServeFile(w, r, path)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")
    if len(email) < 8 || len(password) < 8 {
        http.Error(w, "Invalid emailor password", http.StatusNotAcceptable)
        return
    }
    
    _, err := s.db.GetUser(email)
    if err == nil {
        http.Error(w, "User with \"%v\" email already exists", http.StatusConflict)
        return
    }

    err = s.db.CreateNewUser(email, password)
    if err != nil {
        http.Error(w, "Problem creating new user", http.StatusInternalServerError)
        fmt.Println("Problem creating new user", err)
    }
    return
}

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")
    if len(email) < 8 || len(password) < 8 {
        http.Error(w, "Invalid email or password", http.StatusNotAcceptable)
        return
    }

    user, err := s.db.GetUser(email)
    hashedPassword, _ := util.HashPassword(password)
    saltedHashedPassword, _ := util.HashPassword(fmt.Sprintf("%v%v",hashedPassword, user.Salt))

    if err != nil || saltedHashedPassword != user.SaltedHashedPassword {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
    }

    sessionToken := database.GenerateToken(32)
    csrfToken := database.GenerateToken(32)

    http.SetCookie(w, &http.Cookie{
        Name: "session_token",
        Value: sessionToken,
        Expires: time.Now().Add(24 * time.Hour),
        HttpOnly: true,
    })

    http.SetCookie(w, &http.Cookie{
        Name: "csrf_token",
        Value: csrfToken,
        Expires: time.Now().Add(24 * time.Hour),
        HttpOnly: false,
    })

    // Need to update it in the db as well
    user.SessionToken = sessionToken
    user.CSRFToken= csrfToken

    fmt.Println("Login successfull")

}


func (s *Server) ProtectedHandler(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    if err := s.db.Authorize(r); err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    email := r.FormValue("email") 
    fmt.Fprintf(w, "CSRF Validation successful! Welcome, %s", email) 
}


func (s *Server) LogoutHandler(w http.ResponseWriter, r *http.Request){
    if err := s.db.Authorize(r); err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name: "session_token",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HttpOnly: true,
    })

    http.SetCookie(w, &http.Cookie{
        Name: "csrf_token",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HttpOnly: false,
    })

    email := r.FormValue("email")
    user, _ := s.db.GetUser(email)

    // Need to update it in the db as well
    user.SessionToken = ""
    user.CSRFToken= ""

}



















