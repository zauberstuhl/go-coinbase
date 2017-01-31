package coinbase

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "crypto/sha256"
  "crypto/hmac"
  "fmt"
)

func TestClientFetch(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"Data":{"Epoch":1485347945}}`)
  }))
  defer ts.Close()

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: ts.URL,
  }

  var result APIClientEpoch
  err := a.Fetch("GET", "/v2/time", nil, &result)
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if result.Data.Epoch != 1485347945 {
    t.Error("Expected valid unix timestamp, got ", result.Data.Epoch)
  }
}

func TestClientAuthenticate(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, `{"Data":{"Epoch":1485347945}}`)
  }))
  defer ts.Close()

  req, err := http.NewRequest("GET", "/", nil)
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: ts.URL,
  }
  err = a.Authenticate("/", req, nil)
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }

  key := req.Header.Get("CB-ACCESS-KEY")
  if key != "123" {
    t.Error("Expected key to be 123, got ", key)
  }

  signatureHeader := []byte(req.Header.Get("CB-ACCESS-SIGN"))
  h := hmac.New(sha256.New, []byte("123456"))
  h.Write([]byte("1485347945GET/"))
  signature := fmt.Sprintf("%x", h.Sum(nil))
  if !hmac.Equal(signatureHeader, []byte(signature)) {
    t.Error("Expected equal signatures, got different ones")
  }

  timestamp := req.Header.Get("CB-ACCESS-TIMESTAMP")
  if timestamp != "1485347945" {
    t.Error("Expected valid unix timestamp, got ", timestamp)
  }
}
