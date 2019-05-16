package coinbase

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "fmt"
)

var transactionTmpl = `{
  "data": {
    "id": "8250fe29-f5ef-5fc5-8302-0fbacf6be51e",
    "type": "%s",
    "status": "pending",
    "amount": {
      "amount": "1.00000000",
      "currency": "BTC"
    },
    "native_amount": {
      "amount": "10.00",
      "currency": "USD"
    },
    "description": null,
    "created_at": "2015-03-26T13:42:00-07:00",
    "updated_at": "2015-03-26T15:55:45-07:00",
    "resource": "transaction",
    "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/transactions/8250fe29-f5ef-5fc5-8302-0fbacf6be51e",
    "%s": {
      "id": "5c8216e7-318a-50a5-91aa-2f2cfddfdaab",
      "resource": "%s",
      "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/%ss/5c8216e7-318a-50a5-91aa-2f2cfddfdaab"
    },
    "instant_exchange": false,
    "details": {
      "title": "Bought/Sold bitcoin",
      "subtitle": "using Capital One Bank"
    }
  }
}`

func TestGetTransaction(t *testing.T) {
  var transactionTypes = []string{
    "buy","sell","send","transfer","fiat_deposit",
    "fiat_withdrawal","exchange_deposit",
    "exchange_withdrawal","vault_withdrawal",
  }

  for i, transactionType := range transactionTypes {
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintln(w, fmt.Sprintf(
        transactionTmpl, transactionType, transactionType,
        transactionType, transactionType))
    }))
    defer ts.Close()

    a := APIClient{
      Key: "123",
      Secret: "123456",
      Endpoint: ts.URL,
    }

    transaction, err := a.GetTransaction("", "")
    if err != nil {
      t.Errorf("#%d: Expected nil, got '%+v'", i, err.Error())
    }

    switch transactionType {
      case "buy":
        if transaction.Data.Buy == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "sell":
        if transaction.Data.Sell == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "send":
        if transaction.Data.Send == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "transfer":
        if transaction.Data.Transfer == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "fiat_deposit":
        if transaction.Data.Fiat_deposit == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "fiat_withdrawal":
        if transaction.Data.Fiat_withdrawal == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "exchange_deposit":
        if transaction.Data.Exchange_deposit == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "exchange_withdrawal":
        if transaction.Data.Exchange_withdrawal == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      case "vault_withdrawal":
        if transaction.Data.Vault_withdrawal == nil {
          t.Errorf("#%d: Expected not nil, got nil", i)
        }
      default:
        t.Errorf("#%d: Unknown test type", i)
    }
  }
}
