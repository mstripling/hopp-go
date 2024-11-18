package util

import (
	"crypto/sha256"
	"encoding/hex"
  	"strconv"
  	"fmt"
)


type Payload struct {
	Plain map[string]interface{} `json:"plain"`
	Hash map[string]interface{} `json:"hash"`
}

func Hash(p Payload) (map[string]interface{}, error) {
    result := make(map[string]interface{})
	
    // Extracts keys and unhashed values
    for key, value := range p.Plain {
		result[key] = value
	}
	
	// Extracts keys and hashes values before adding them to result
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
        result[key] = hex.EncodeToString(hash[:]) 
  }

  return result, nil
}
