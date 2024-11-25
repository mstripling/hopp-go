package util

import(
  "net/http"
  "errors"
)

func Initialize(w http.ResponseWriter, r *http.Request, v *RawPayload) error {
    if v.Endpoint == ""{
      return errors.New("No endpoint")
    }
    if v.Plain == nil {
      v.Plain = make(map[string]interface{})
    }
    
    if v.Hash == nil {
      v.Hash = make(map[string]interface{})
    }
    return nil
}
