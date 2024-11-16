pachage hash

import (
	"crypto"
	"encoding/json"
	"fmt"
)

type Payload struct {
	Plain map[string]interface{} `json:"plain"`
	Hash map[string]interface{} `json:"hash"`
}

func Hash(s string){
	
}

