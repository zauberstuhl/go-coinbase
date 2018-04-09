package coinbase

import (
  "testing"
)

func TestBuys(t *testing.T) {
  s := serverHelper(`
{
  "data": {
    "id": "9e14d574-30fa-5d85-b02c-6be0d851d61d",
    "status": "created",
    "payment_method": {
      "id": "83562370-3e5c-51db-87da-752af5ab9559",
      "resource": "payment_method",
      "resource_path": "/v2/payment-methods/83562370-3e5c-51db-87da-752af5ab9559"
    },
    "transaction": {
      "id": "4117f7d6-5694-5b36-bc8f-847509850ea4",
      "resource": "transaction",
      "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/transactions/441b9494-b3f0-5b98-b9b0-4d82c21c252a"
    },
    "amount": {
      "amount": "10.00000000",
      "currency": "BTC"
    },
    "total": {
      "amount": "102.01",
      "currency": "USD"
    },
    "subtotal": {
      "amount": "101.00",
      "currency": "USD"
    },
    "created_at": "2015-03-26T23:43:59-07:00",
    "updated_at": "2015-03-26T23:44:09-07:00",
    "resource": "buy",
    "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/buys/9e14d574-30fa-5d85-b02c-6be0d851d61d",
    "committed": true,
    "instant": false,
    "fee": {
      "amount": "1.01",
      "currency": "USD"
    },
    "payout_at": "2015-04-01T23:43:59-07:00"
  }
}
  `)
  defer s.Close()

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: s.URL,
  }

  buys, err := a.ShowBuy("123", "456")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if buys.Data.Id != "9e14d574-30fa-5d85-b02c-6be0d851d61d" {
    t.Error("Expected 9e14d574-30fa-5d85-b02c-6be0d851d61d, got ",
      buys.Data.Id)
  }

  wrapper := func(data interface{}) {
    switch r := data.(type) {
    case APIResource:
    case APIBalance:
    default:
      t.Error("Expected APIBalance or APIResource, got ", r)
    }
  }
  wrapper(buys.Data.Payment_method)
  wrapper(buys.Data.Transaction)
  wrapper(buys.Data.Amount)
  wrapper(buys.Data.Total)
  wrapper(buys.Data.Subtotal)

  if buys.Data.Created_at.Year() != 2015 {
    t.Error("Expected 2015, got ", buys.Data.Created_at.Year())
  }
  if buys.Data.Updated_at.Year() != 2015 {
    t.Error("Expected 2015, got ", buys.Data.Created_at.Year())
  }
}
