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
