package coinbase

import (
  "testing"
  "time"
)

func TestGetBuyPrice(t *testing.T) {
  a := APIClient{}
  balance, err := a.GetBuyPrice(
    ConfigPrice{
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
  a := APIClient{}
  balance, err := a.GetSellPrice(
    ConfigPrice{
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
  a := APIClient{}
  date, err := time.Parse("2006-01-02", "2017-01-01")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }

  balance, err := a.GetSpotPrice(
    ConfigPrice{
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
  if balance.Data.Amount <= 0.0 {
    t.Error("Expected amount > 0.0, got ", balance.Data.Amount)
  }
}

