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

Example:
 2bbf394c-193b-5b2a-9155-3b4732659ede

*/
type BuyId string

// ListBuys requires an account ID and returns an APIWalletTransferList struct
func (a *APIClient) ListBuys(id AccountId) (buys APIWalletTransferList, err error) {
  path := pathHelper("/v2/accounts/%s/buys", id)
  err = a.Fetch("GET", path, nil, &buys)
  if err != nil {
    return
  }
  return
}

// ShowBuy requires an account ID, buy ID and returns an APIWalletTransfer struct
func (a *APIClient) ShowBuy(id AccountId, bid BuyId) (buys APIWalletTransfer, err error) {
  path := pathHelper("/v2/accounts/%s/buys/%s", id, bid)
  err = a.Fetch("GET", path, nil, &buys)
  if err != nil {
    return
  }
  return
}

// PlaceBuyOrder requires an account ID, APIWalletTransferOrder and returns an APIWalletTransfer struct
func (a *APIClient) PlaceBuyOrder(id AccountId, order APIWalletTransferOrder) (buys APIWalletTransfer, err error) {
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

// CommitBuy requires an account ID, buy ID and returns an APIWalletTransfer struct
func (a *APIClient) CommitBuy(id AccountId, bid BuyId) (buys APIWalletTransfer, err error) {
  path := pathHelper("/v2/accounts/%s/buys/%s/commit", id, bid)
  err = a.Fetch("POST", path, nil, &buys)
  if err != nil {
    return
  }
  return
}
