package coinbase

import (
  "testing"
)

func TestAccounts(t *testing.T) {
  s := serverHelper(`
 {
   "data": {
       "id": "2bbf394c-193b-5b2a-9155-3b4732659ede",
       "name": "My Wallet",
       "primary": true,
       "type": "wallet",
       "currency": "BTC",
       "balance": {
           "amount": "39.59000000",
           "currency": "BTC"
       },
       "native_balance": {
           "amount": "395.90",
           "currency": "USD"
       },
       "created_at": "2015-01-31T20:49:02Z",
       "updated_at": "2015-01-31T20:49:02Z",
       "resource": "account",
       "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede"
   }
 }
  `)
  defer s.Close()

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: s.URL,
  }

  acc, err := a.Account("2bbf394c-193b-5b2a-9155-3b4732659ede")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if acc.Data.Id != "2bbf394c-193b-5b2a-9155-3b4732659ede" {
    t.Error("Expected 2bbf394c-193b-5b2a-9155-3b4732659ede, got ", acc.Data.Id)
  }
  if acc.Data.Balance.Amount != 39.59 {
    t.Error("Expected 39.59, got ", acc.Data.Balance.Amount)
  }
  if acc.Data.Native_balance.Amount != 395.90 {
    t.Error("Expected 395.90, got ", acc.Data.Native_balance.Amount)
  }
  if acc.Data.Created_at.Year() != 2015 {
    t.Error("Expected 2015, got ", acc.Data.Created_at.Year())
  }
  if acc.Data.Updated_at.Year() != 2015 {
    t.Error("Expected 2015, got ", acc.Data.Created_at.Year())
  }
}
