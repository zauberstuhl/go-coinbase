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

/*

Example Response:

 {
   "id": "57ffb4ae-0c59-5430-bcd3-3f98f797a66c",
   "type": "send",
   "status": "completed",
   "amount": {
     "amount": "-0.00100000",
     "currency": "BTC"
   },
   "native_amount": {
     "amount": "-0.01",
     "currency": "USD"
   },
   "description": null,
   "created_at": "2015-03-11T13:13:35-07:00",
   "updated_at": "2015-03-26T15:55:43-07:00",
   "resource": "transaction",
   "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/transactions/57ffb4ae-0c59-5430-bcd3-3f98f797a66c",
   "network": {
     "status": "off_blockchain",
     "name": "bitcoin"
   },
   "to": {
     "id": "a6b4c2df-a62c-5d68-822a-dd4e2102e703",
     "resource": "user",
     "resource_path": "/v2/users/a6b4c2df-a62c-5d68-822a-dd4e2102e703"
   },
   "instant_exchange": false,
   "details": {
     "title": "Sent bitcoin",
     "subtitle": "to User 2"
   }
 }

*/
type APITransactionsDataDetails struct {
  Title string
  Subtitle string
}
type APITransactionsDataNetwork struct {
  Status string
  Name string
}
type APITransactionsData  struct {
  Id string
  Type string
  Status string
  Amount APIBalance
  Native_amount APIBalance
  Description string
  Created_at string
  Updated_at string
  Resource string
  Resource_path string
  Network APITransactionsDataNetwork
  To APIResource
  Buy *APIResource `json:"buy,omitempty"`
  Sell *APIResource `json:"sell,omitempty"`
  Send *APIResource `json:"send,omitempty"`
  Request *APIResource `json:"request,omitempty"`
  Transfer *APIResource `json:"transfer,omitempty"`
  Fiat_deposit *APIResource `json:"fiat_deposit,omitempty"`
  Fiat_withdrawal *APIResource `json:"fiat_withdrawal,omitempty"`
  Exchange_deposit *APIResource `json:"exchange_deposit,omitempty"`
  Exchange_withdrawal *APIResource `json:"exchange_withdrawal,omitempty"`
  Vault_withdrawal *APIResource `json:"vault_withdrawal,omitempty"`
  Instant_exchange bool
  Details APITransactionsDataDetails
}
type APITransactions struct {
  Pagination APIPagination
  Data []APITransactionsData
  Errors []Error
}
type APITransaction struct {
  Data APITransactionsData
  Errors []Error
}
// GetTransactions requires an account ID and returns an APITransactions struct
func (a *APIClient) GetTransactions(id string) (trans APITransactions, err error) {
  err = a.Fetch("GET", "/v2/accounts/" + id + "/transactions", nil, &trans)
  if err != nil {
    return
  }
  return
}

// GetTransaction requires an account ID, transaction ID and returns an APITransaction struct
func (a *APIClient) GetTransaction(id, transId string) (trans APITransaction, err error) {
  path := "/v2/accounts/" + id + "/transactions/" + transId
  err = a.Fetch("GET", path, nil, &trans)
  if err != nil {
    return
  }
  return
}

type APITransactionsSend struct {
  Type string `json:"type"`
  To string `json:"to"`
  Amount float64 `json:"amount"`
  Currency string `json:"currency"`
  Fee string `json:"fee"`
}


// SendTransferRequestMoney requires an account ID, APITransactionsSend struct
// and returns am APITransaction struct
// TODO move account ID into the APITransactionsSend struct
func (a *APIClient) SendTransferRequestMoney(id string, send APITransactionsSend) (trans APITransaction, err error) {
  // TODO implement idem
  err = a.Fetch("POST", "/v2/accounts/" + id + "/transactions", send, &trans)
  if err != nil {
    return
  }
  return
}

// CompleteRequestMoney requires an account ID, transaction ID
func (a *APIClient) CompleteRequestMoney(id, transId string) (err error) {
  path := "/v2/accounts/" + id + "/transactions/" + transId + "/complete"
  err = a.Fetch("GET", path, nil, nil)
  if err != nil {
    return
  }
  return
}

// ResendRequestMoney requires an account ID, transaction ID
func (a *APIClient) ResendRequestMoney(id, transId string) (err error) {
  path := "/v2/accounts/" + id + "/transactions/" + transId + "/resend"
  err = a.Fetch("GET", path, nil, nil)
  if err != nil {
    return
  }
  return
}

// CancelRequestMoney requires an account ID, transaction ID
func (a *APIClient) CancelRequestMoney(id, transId string) (err error) {
  path := "/v2/accounts/" + id + "/transactions/" + transId
  err = a.Fetch("DELETE", path, nil, nil)
  if err != nil {
    return
  }
  return
}
