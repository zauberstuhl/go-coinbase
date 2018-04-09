package coinbase

import (
  "testing"
)

func TestGetCurrencies(t *testing.T) {
  a := APIClient{
    Key: "123",
    Secret: "123456",
  }
  list, err := a.GetCurrencies()
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if len(list.Data) <= 0 {
    t.Error("Expected list to be greater then zero")
  }
  for i, curr := range list.Data {
    if len(curr.Id) != 3 {
      t.Error(i, ". Expected id length equal three")
    }
    if len(curr.Name) <= 0 {
      t.Error(i, ". Expected name length to be greater then zero")
    }
    if curr.Min_size <= 0.0 {
      t.Error(i, ". Expected min_size to be greater then zero")
    }
  }
}
