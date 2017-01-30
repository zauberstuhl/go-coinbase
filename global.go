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
  "fmt"
  "time"
)

// AccountId is a unique ID for your wallet
//
// Example:
//  2bbf394c-193b-5b2a-9155-3b4732659ede
type AccountId string

//// API Structs /////////////////

type Error struct {
  Id string `json:"id"`
  Message string `json:"message"`
}

type APIName struct {
  Name string `json:"name"`
}

/*

Example Response:

 "pagination": {
   "ending_before": null,
   "starting_after": null,
   "limit": 25,
   "order": "desc",
   "previous_uri": null,
   "next_uri": null
 },

*/
type APIPagination struct {
  Ending_before string
  Starting_after string
  Limit int
  Order string
  Previous_uri string
  Next_uri string
}

/*

Example Response:

 "balance": {
     "amount": "39.59000000",
     "currency": "BTC"
 },

*/
type APIBalance struct {
  Amount float64 `json:",string"`
  Currency string
}

/*

Example Response:

 "to": {
   "id": "a6b4c2df-a62c-5d68-822a-dd4e2102e703",
   "resource": "user",
   "resource_path": "/v2/users/a6b4c2df-a62c-5d68-822a-dd4e2102e703"
 },

*/
type APIResource struct {
  Id string
  Resource string
  Resource_path string
}

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
type APIWalletTransferData struct {
  Id string
  Status string
  Payment_method APIResource
  Transaction APIResource
  Amount APIBalance
  Total APIBalance
  Subtotal APIBalance
  Created_at *time.Time
  Updated_at *time.Time
  Resource string
  Resource_path string
  Committed bool
  Instant bool
  Fee APIBalance
  Payout_at string
}
type APIWalletTransferList struct {
  Pagination APIPagination
  Data []APIWalletTransferData
  Errors []Error
}
type APIWalletTransfer struct {
  Data APIWalletTransferData
  Errors []Error
}

/*

Example request:

 {
   "amount": "10",
   "currency": "BTC",
   "payment_method": "83562370-3e5c-51db-87da-752af5ab9559"
 }

*/
type APIWalletTransferOrder struct {
  Amount *float64 `json:"amount"`
  Total *float64 `json:"total"`
  Currency string `json:"currency"`
  Payment_method string `json:"payment_method"`
  Agree_btc_amount_varies bool `json:"agree_btc_amount_varies"`
  Commit bool `json:"commit"`
  Quote bool `json:"quote"`
}

//// Configuration Structs //

type ConfigPrice struct{
  From string
  To string
  Date time.Time
}

//// Helper Functions ///////

// pathHelper is a simple wrapper for Sprintf
// and exists only to reduce import expressions
func pathHelper(f string, p... interface{}) string {
  return fmt.Sprintf(f, p...)
}
