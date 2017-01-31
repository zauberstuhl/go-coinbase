package coinbase

import (
  "net/http"
  "net/http/httptest"
  "fmt"
)

func serverHelper(body string) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, body)
  }))
}
