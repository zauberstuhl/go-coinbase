package coinbase_test

import (
  "testing"
  "zauberstuhl/coinbase"
)

func TestGetCurrentTime(t *testing.T) {
  a := coinbase.APIClient{
    Key: "123",
    Secret: "123456",
  }
  time, err := a.GetCurrentTime()
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if time.Data.Epoch < 1485347945 {
    t.Error("Expected valid unix timestamp, got ", time.Data.Epoch)
  }
}
