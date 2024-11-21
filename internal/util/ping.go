package util

import(
  "net/http"
)

func Ping(r *http.Request, p map[string]interface{}) (*http.Response, error) {
  req, err := http.NewRequest(r.Method, "http://localhost:8080/buyer", nil)
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
