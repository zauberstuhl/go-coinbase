/*
 * Coinbase Golang API Library
 *
 * Copyright (C) 2017 Lukas Matt <lukas@zauberstuhl.de>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package coinbase

import (
  "bytes"
  "encoding/json"
)

/*

Example Response:

 {
   "id": "67e0eaec-07d7-54c4-a72c-2e92826897df",
   "status": "completed",
   "payment_method": {
     "id": "83562370-3e5c-51db-87da-752af5ab9559",
     "resource": "payment_method",
     "resource_path": "/v2/payment-methods/83562370-3e5c-51db-87da-752af5ab9559"
   },
   "transaction": {
     "id": "441b9494-b3f0-5b98-b9b0-4d82c21c252a",
     "resource": "transaction",
     "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/transactions/441b9494-b3f0-5b98-b9b0-4d82c21c252a"
   },
   "amount": {
     "amount": "1.00000000",
     "currency": "BTC"
   },
   "total": {
     "amount": "10.25",
     "currency": "USD"
   },
   "subtotal": {
     "amount": "10.10",
     "currency": "USD"
   },
   "created_at": "2015-01-31T20:49:02Z",
   "updated_at": "2015-02-11T16:54:02-08:00",
   "resource": "buy",
   "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/buys/67e0eaec-07d7-54c4-a72c-2e92826897df",
   "committed": true,
   "instant": false,
   "fee": {
     "amount": "0.15",
     "currency": "USD"
   },
   "payout_at": "2015-02-18T16:54:00-08:00"
 }

*/
type APIBuysSellsData struct {
  Id string
  Status string
  Payment_method APIResource
  Transaction APIResource
  Amount APIBalance
  Total APIBalance
  Subtotal APIBalance
  Created_at string
  Updated_at string
  Resource string
  Resource_path string
  Committed bool
  Instant bool
  Fee APIBalance
  Payout_at string
}
type APIBuysSellsList struct {
  Pagination APIPagination
  Data []APIBuysSellsData
}
type APIBuysSells struct {
  Data APIBuysSellsData
}
type BuyId string
// ListBuys requires an account ID and returns an APIBuysSellsList struct
func (a *APIClient) ListBuys(id AccountId) (buys APIBuysSellsList, err error) {
  path := pathHelper("/v2/accounts/%s/buys", id)
  err = a.Fetch("GET", path, nil, &buys)
  if err != nil {
    return
  }
  return
}

// ShowBuy requires an account ID, buy ID and returns an APIBuysSells struct
func (a *APIClient) ShowBuy(id AccountId, bid BuyId) (buys APIBuysSells, err error) {
  path := pathHelper("/v2/accounts/%s/buys/%s", id, bid)
  err = a.Fetch("GET", path, nil, &buys)
  if err != nil {
    return
  }
  return
}

/*

Example request:

 {
   "amount": "10",
   "currency": "BTC",
   "payment_method": "83562370-3e5c-51db-87da-752af5ab9559"
 }

*/
type APIBuysBuySellOrder struct {
  Amount float64
  Total float64
  Currency string
  Payment_method string
  Agree_btc_amount_varies bool
  Commit bool
  Quote bool
}
// PlaceBuyOrder requires an account ID, APIBuysBuySellOrder and returns an APIBuysSells struct
func (a *APIClient) PlaceBuyOrder(id AccountId, order APIBuysBuySellOrder) (buys APIBuysSells, err error) {
  data, err := json.Marshal(order)
  if err != nil {
    return buys, err
  }
  path := pathHelper("/v2/accounts/%s/buys", id)
  err = a.Fetch("POST", path, bytes.NewBuffer([]byte(data)), &buys)
  if err != nil {
    return
  }
  return
}

// CommitBuy requires an account ID, buy ID and returns an APIBuysSells struct
func (a *APIClient) CommitBuy(id AccountId, bid BuyId) (buys APIBuysSells, err error) {
  path := pathHelper("/v2/accounts/%s/buys/%s/commit", id, bid)
  err = a.Fetch("POST", path, nil, &buys)
  if err != nil {
    return
  }
  return
}
