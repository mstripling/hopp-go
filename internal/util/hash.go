package util

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "fmt"
    "math/rand"
    "time"
)


type RawPayload struct {
    Plain map[string]interface{}  `json:"plain"`
    Hash map[string]interface{}   `json:"hash"`
    Endpoint string               `json:"endpoint"`
    Test bool                     `json:"test,omitempty"`
}


func TransformAndFormat(p *RawPayload) (map[string]interface{}, error) {
    pingBody := make(map[string]interface{})

    for key, value := range p.Plain {
        pingBody[key] = value
    }

    for key, value := range p.Hash {
        var strValue string
        
        // Type assertion to string
        switch v := value.(type) {
        case string:
            strValue = v
        case int:
            strValue = strconv.Itoa(v)
        case float64:
            strValue = fmt.Sprintf("%f", v)
        default:
            return nil, fmt.Errorf("unsupported value type for key '%s': %T", key, value)
        }

        hash := sha256.Sum256([]byte(strValue))
        pingBody[key] = hex.EncodeToString(hash[:]) 
    }

    return pingBody, nil
}

func HashPassword (p string) (string, string) {
    // Create seed for salt 
    rand.Seed(time.Now().UnixNano())

    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:',.<>?/`~\""
    length := 16
    var saltByte []byte
    for i:=0; i < length; i++ {
        randomIndex := rand.Intn(len(charset))
        saltByte = append(saltByte, charset[randomIndex])
    }
    salt := string(saltByte[:])

    hashedPasswordBytes := sha256.Sum256([]byte(p))
    hashedPassword := hex.EncodeToString(hashedPasswordBytes[:]) 

    return hashedPassword, salt
}














