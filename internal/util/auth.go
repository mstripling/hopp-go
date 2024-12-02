package util

import(
    "crypto/rand"
    "encoding/base64"
    "log"
    "crypto/sha256"
)

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
