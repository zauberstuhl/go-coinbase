package coinbase_test

import (
  "testing"
  "zauberstuhl/coinbase"
  "time"
)

func TestGetBuyPrice(t *testing.T) {
  a := coinbase.APIClient{}
  balance, err := a.GetBuyPrice(
    coinbase.ConfigPrice{
      From: "BTC",
      To: "EUR",
  })
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if balance.Data.Currency != "EUR" {
    t.Error("Expected EUR, got ", balance.Data.Currency)
  }
  if balance.Data.Amount <= 0.0 {
    t.Error("Expected <= 0.0, got ", balance.Data.Amount)
  }
}

func TestGetSellPrice(t *testing.T) {
  a := coinbase.APIClient{}
  balance, err := a.GetSellPrice(
    coinbase.ConfigPrice{
      From: "BTC",
      To: "EUR",
  })
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if balance.Data.Currency != "EUR" {
    t.Error("Expected EUR, got ", balance.Data.Currency)
  }
  if balance.Data.Amount <= 0.0 {
    t.Error("Expected <= 0.0, got ", balance.Data.Amount)
  }
}

func TestGetSpotPrice(t *testing.T) {
  a := coinbase.APIClient{}
  date, err := time.Parse("2006-01-02", "2017-01-01")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }

  balance, err := a.GetSpotPrice(
    coinbase.ConfigPrice{
      From: "BTC",
      To: "EUR",
      Date: date,
  })
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if balance.Data.Currency != "EUR" {
    t.Error("Expected EUR, got ", balance.Data.Currency)
  }
  if balance.Data.Amount < 500.0 {
    t.Error("Expected amount >= 800, got ", balance.Data.Amount)
  }
}

