package coinbase

import (
  "net/http"
  "net/http/httptest"
  "fmt"
  "testing"
)

func serverHelper(t *testing.T, method, uri, body *string) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.RequestURI == "/v2/time" {
      // NOTE this is used on a signing request
      fmt.Fprintln(w, `{
        "data": {
          "iso": "2015-06-23T18:02:51Z",
          "epoch": 1435082571
        }
      }`)
      return
    }

    if *method != r.Method {
      t.Errorf("Expected http method %s, got %s", *method, r.Method)
    }

    if *uri != r.RequestURI {
      t.Errorf("Expected request uri %s, got %s", *uri, r.RequestURI)
    }

    fmt.Fprintln(w, *body)
  }))
}
