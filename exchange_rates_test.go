package coinbase

import (
  "testing"
)

func TestGetExchangeRates(t *testing.T) {
  a := APIClient{
    Key: "123",
    Secret: "123456",
  }
  exch, err := a.GetExchangeRates("BTC")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if len(exch.Data.Currency) != 3 {
    t.Error("Expected currency length equal three")
  }
  for key, value := range exch.Data.Rates {
    if len(key) != 3 {
      t.Error("Expected ", key, " length equal three")
    }
    rate, err := value.Float64()
    if err != nil {
      t.Error("Expected nil, got ", err.Error())
    }
    if rate <= 0.0 {
      t.Error("Expected value to be greater then zero")
    }
  }
}
