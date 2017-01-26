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
type DepositId string

// ListDeposits requires an account ID and returns an APIWalletTransferList struct
func (a *APIClient) ListDeposits(id AccountId) (deposits APIWalletTransferList, err error) {
  path := pathHelper("/v2/accounts/%s/deposits", id)
  err = a.Fetch("GET", path, nil, &deposits)
  if err != nil {
    return
  }
  return
}

// ShowDeposit requires an account ID, deposit ID and returns an APIWalletTransfer struct
func (a *APIClient) ShowDeposit(id AccountId, did DepositId) (deposits APIWalletTransfer, err error) {
  path := pathHelper("/v2/accounts/%s/deposits/%s", id, did)
  err = a.Fetch("GET", path, nil, &deposits)
  if err != nil {
    return
  }
  return
}

// PlaceDepositOrder requires an account ID, APIWalletTransferOrder and returns an APIWalletTransfer struct
func (a *APIClient) PlaceDepositOrder(id AccountId, order APIWalletTransferOrder) (deposits APIWalletTransfer, err error) {
  data, err := json.Marshal(order)
  if err != nil {
    return deposits, err
  }
  path := pathHelper("/v2/accounts/%s/deposits", id)
  err = a.Fetch("POST", path, bytes.NewBuffer([]byte(data)), &deposits)
  if err != nil {
    return
  }
  return
}

// CommitDeposit requires an account ID, deposit ID and returns an APIWalletTransfer struct
func (a *APIClient) CommitDeposit(id AccountId, did DepositId) (deposits APIWalletTransfer, err error) {
  path := pathHelper("/v2/accounts/%s/buys/%s/commit", id, did)
  err = a.Fetch("POST", path, nil, &deposits)
  if err != nil {
    return
  }
  return
}
