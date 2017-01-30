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
   "id": "dd3183eb-af1d-5f5d-a90d-cbff946435ff",
   "address": "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa",
   "name": "One off payment",
   "created_at": "2015-01-31T20:49:02Z",
   "updated_at": "2015-03-31T17:25:29-07:00",
   "network": "bitcoin",
   "resource": "address",
   "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/addresses/dd3183eb-af1d-5f5d-a90d-cbff946435ff"
 }

*/
type APIAddressData struct {
  Id string
  Address string
  Name string
  Created_at *time.Time
  Updated_at *time.Time
  Network string
  Resource string
  Resource_path string
}
type APIAddresses struct {
  Pagination APIPagination
  Data []APIAddressData
  Errors []Error
}
type APIAddress struct {
  Data APIAddressData
  Errors []Error
}
// GetAddresses requires an address ID and returns an APIAddresses struct
func (a *APIClient) GetAddresses(id string) (addresses APIAddresses, err error) {
  err = a.Fetch("GET", "/v2/accounts/" + id + "/addresses", nil, &addresses)
  if err != nil {
    return
  }
  return
}

// GetAddress requires an account ID and address ID. It will return an APIAddress struct
func (a *APIClient) GetAddress(id, addressId string) (address APIAddress, err error) {
  err = a.Fetch("GET", "/v2/accounts/" + id + "/addresses/" + addressId, nil, &address)
  if err != nil {
    return
  }
  return
}

// ListAddressTransactions requires an account ID and address ID. It will return an APITransactions struct
func (a *APIClient) ListAddressTransactions(id, address string) (trans APITransactions, err error) {
  path := "/v2/accounts/" + id + "/addresses/" + address + "/transactions"
  err = a.Fetch("GET", path, nil, &trans)
  if err != nil {
    return
  }
  return
}

// CreateAddress requires an account ID
// and the address Name as parameter.
// It will return an APIAddress struct
func (a *APIClient) CreateAddress(id, name string) (address APIAddress, err error) {
  var body APIName
  body.Name = name
  err = a.Fetch("POST", "/v2/accounts/" + id + "/addresses", body, &address)
  if err != nil {
    return
  }
  return
}
