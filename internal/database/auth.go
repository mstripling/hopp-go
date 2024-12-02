package database

import(
    "crypto/rand"
    "encoding/base64"
    "log"
    "errors"
    "net/http"
    //"crypto/sha256"

)

var AuthError = errors.New("Unauthorized")

type User struct {
    Email string
    SaltedHashedPassword string
    Salt string
    SessionToken string
    CSRFToken string
}

func GenerateToken(length int) string {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        log.Fatalf("Failed to generate token: %v", err)
    }
    
    return base64.URLEncoding.EncodeToString(bytes)
}

func (s *service) Authorize(r *http.Request) error {
    email := r.FormValue("email")

    user, err := s.GetUser(email)
    if err != nil {
        return AuthError 
    }

    st, err := r.Cookie("session_token")
    if err != nil || st.Value == "" || st.Value != user.SessionToken {
        return AuthError
    }

    csrf := r.Header.Get("X-CSRF-Token")
    if csrf == "" || csrf != user.CSRFToken {
        return AuthError
    }

    return nil
}
