package coinbase

import (
  "testing"
)

func TestAccount(t *testing.T) {
  method := "GET"
  uri := "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede"
  body := `{
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
  }`
  s := serverHelper(t, &method, &uri, &body)
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
  account(t, acc.Data)

  method = "POST"
  uri = "/v2/accounts"
  acc, err = a.CreateAccount("My Wallet")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  account(t, acc.Data)

  uri = "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/primary"
  acc, err = a.SetPrimaryAccount("2bbf394c-193b-5b2a-9155-3b4732659ede")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  account(t, acc.Data)

  method = "PUT"
  uri = "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede"
  acc, err = a.UpdateAccount("2bbf394c-193b-5b2a-9155-3b4732659ede", "My Wallet")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  account(t, acc.Data)

  method = "DELETE"
  acc, err = a.DeleteAccount("2bbf394c-193b-5b2a-9155-3b4732659ede")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  account(t, acc.Data)
}

func TestAccounts(t *testing.T) {
  method := "GET"
  uri := "/v2/accounts"
  body := `{
    "pagination": {
      "ending_before": null,
      "starting_after": null,
      "limit": 25,
      "order": "desc",
      "previous_uri": null,
      "next_uri": null
    },
    "data": [
      {
        "id": "2bbf394c-193b-5b2a-9155-3b4732659ede",
        "name": "My Vault",
        "primary": false,
        "type": "vault",
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
        "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede",
        "ready": true
      }
    ]
  }`
  s := serverHelper(t, &method, &uri, &body)
  defer s.Close()

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: s.URL,
  }

  accs, err := a.Accounts()
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }

  if len(accs.Data) != 1 {
    t.Errorf("Expected len to be 1, got %d", len(accs.Data))
  }

  for _, data := range accs.Data {
    account(t, data)
  }
}

func account(t *testing.T, acc APIAccountData) {
  if acc.Id != "2bbf394c-193b-5b2a-9155-3b4732659ede" {
    t.Error("Expected 2bbf394c-193b-5b2a-9155-3b4732659ede, got ", acc.Id)
  }
  if acc.Balance.Amount != 39.59 {
    t.Error("Expected 39.59, got ", acc.Balance.Amount)
  }
  if acc.Native_balance.Amount != 395.90 {
    t.Error("Expected 395.90, got ", acc.Native_balance.Amount)
  }
  if acc.Created_at.Year() != 2015 {
    t.Error("Expected 2015, got ", acc.Created_at.Year())
  }
  if acc.Updated_at.Year() != 2015 {
    t.Error("Expected 2015, got ", acc.Created_at.Year())
  }
}
