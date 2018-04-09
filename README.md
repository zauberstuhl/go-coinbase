# Coinbase Golang API Library

[![Build Status](https://travis-ci.org/Zauberstuhl/go-coinbase.svg?branch=master)](https://travis-ci.org/Zauberstuhl/go-coinbase) 
[![GoDoc](https://godoc.org/github.com/Zauberstuhl/go-coinbase?status.svg)](http://godoc.org/github.com/Zauberstuhl/go-coinbase)
[![Codecov](https://codecov.io/gh/Zauberstuhl/go-coinbase/branch/master/graph/badge.svg)](https://codecov.io/gh/Zauberstuhl/go-coinbase)

The library was tested against coinbase.com APIv2

## Installation

    go get github.com/Zauberstuhl/go-coinbase

    # or use gopkg for specific versions
    go get gopkg.in/Zauberstuhl/go-coinbase.v1.0.0

## Supported API Calls

* Wallet Endpoints
 - [x] Users
 - [x] Accounts
 - [x] Addresses
 - [x] Transactions
 - [x] Buys
 - [x] Sells
 - [x] Deposits
 - [x] Withdrawals
 - [x] Payment methods
* Data Endpoints
 - [x] Currencies
 - [x] Exchange rates
 - [x] Prices
 - [x] Time

## Example

    import "github.com/Zauberstuhl/go-coinbase"

    c := coinbase.APIClient{
      Key: "123",
      Secret: "123456",
    }

    acc, err := c.Accounts()
    if err != nil {
      fmt.Println(err)
      return
    }

    for i, acc := range accounts.Data {
      fmt.Printf("ID: %s\nName: %s\nType: %s\nAmount: %f\nCurrency: %s\n",
        acc.Id, acc.Name, acc.Type,
        acc.Balance.Amount, acc.Balance.Currency)
    }

    # sample output
    ID: 1234-12-1234-1232
    Name: Test Wallet
    Type: BTC
    Amount: 0.0
    Currency: EUR
    [...]

## Unit Tests

Run all available unit tests via:

    go test

Most of the tests require a internet connection!
