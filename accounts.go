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

import "time"

/*

Example Response:

 {
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

*/
type APIAccountData struct {
  Id string
  Name string
  Primary bool
  Type string
  Currency string
  Balance APIBalance
  Native_balance APIBalance
  Created_at *time.Time
  Updated_at *time.Time
  Resource string
  Resource_path string
}
type APIAccount struct {
  Data APIAccountData
  Errors []Error
}
// Account requires a account ID and returns and APIAccount struct
func (a *APIClient) Account(id string) (account APIAccount, err error) {
  err = a.Fetch("GET", "/v2/accounts/" + id, nil, &account)
  if err != nil {
    return
  }
  return
}

/*

Example Response:

 {
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
       "id": "58542935-67b5-56e1-a3f9-42686e07fa40",
       "name": "My Vault",
       "primary": false,
       "type": "vault",
       "currency": "BTC",
       "balance": {
         "amount": "4.00000000",
         "currency": "BTC"
       },
       "native_balance": {
         "amount": "40.00",
         "currency": "USD"
       },
       "created_at": "2015-01-31T20:49:02Z",
       "updated_at": "2015-01-31T20:49:02Z",
       "resource": "account",
       "resource_path": "/v2/accounts/58542935-67b5-56e1-a3f9-42686e07fa40",
       "ready": true
     },
 [...]

*/
type APIAccounts struct {
  Pagination APIPagination
  Data []APIAccountData
  Errors []Error
}

// Accounts returns all known accounts as APIAccounts struct
func (a *APIClient) Accounts() (accounts APIAccounts, err error) {
  err = a.Fetch("GET", "/v2/accounts", nil, &accounts)
  if err != nil {
    return
  }
  return
}

// CreateAccount requires an account name as parameter and returns an APIAccount struct
func (a *APIClient) CreateAccount(name string) (account APIAccount, err error) {
  var body APIName
  body.Name = name
  err = a.Fetch("POST", "/v2/accounts", body, &account)
  if err != nil {
    return
  }
  return
}

// SetPrimaryAccount requires an account ID and returns an APIAccount struct
func (a *APIClient) SetPrimaryAccount(id string) (account APIAccount, err error) {
  path := "/v2/accounts/" + id + "/primary"
  err = a.Fetch("POST", path, nil, &account)
  if err != nil {
    return
  }
  return
}

// UpdateAccount requires an account ID, Name
// as parameter and returns an APIAccount struct
func (a *APIClient) UpdateAccount(id, name string) (account APIAccount, err error) {
  path := "/v2/accounts/" + id
  var body APIName
  body.Name = name
  err = a.Fetch("PUT", path, body, &account)
  if err != nil {
    return
  }
  return
}

// DeleteAccount requires an account ID and returns an APIAccount struct
func (a *APIClient) DeleteAccount(id string) (account APIAccount, err error) {
  path := "/v2/accounts/" + id
  err = a.Fetch("DELETE", path, nil, &account)
  if err != nil {
    return
  }
  return
}
