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

var tests = []struct{
  Func http.HandlerFunc
  SellNil bool
  BuyNil bool
}{
  {
    Func: func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintln(w, fmt.Sprintf(transactionTmpl, "buy", "buy", "buy", "buy"))
    },
    SellNil: true,
    BuyNil: false,
  },
  {
    Func: func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintln(w, fmt.Sprintf(transactionTmpl, "sell", "sell", "sell", "sell"))
    },
    SellNil: false,
    BuyNil: true,
  },
  {
    Func: func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprintln(w, transactionTmpl)
    },
    SellNil: true,
    BuyNil: true,
  },
}

func TestGetTransaction(t *testing.T) {
  for i, test := range tests {
    ts := httptest.NewServer(test.Func)
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

    if !test.SellNil && transaction.Data.Sell == nil {
      t.Errorf("#%d: Expected not nil, got nil", i)
    }

    if test.SellNil && transaction.Data.Sell != nil {
      t.Errorf("#%d: Expected nil, got '%+v'", i, *transaction.Data.Sell)
    }

    if !test.BuyNil && transaction.Data.Buy == nil {
      t.Errorf("#%d: Expected not nil, got nil", i)
    }

    if test.BuyNil && transaction.Data.Buy != nil {
      t.Errorf("#%d: Expected nil, got '%+v'", i, *transaction.Data.Buy)
    }
  }
}
