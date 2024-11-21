package util

import(
  "net/http"
  "bytes"
  "encoding/json"
)

func Ping(r *http.Request, p map[string]interface{}, e string) (*http.Response, error) {
	// Convert the map to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}  
  
  req, err := http.NewRequest(r.Method, e, bytes.NewReader(jsonData))
  if err != nil {
    return nil, err
  }
  
	// Copy all headers from r to req
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }
  return resp, nil
}
