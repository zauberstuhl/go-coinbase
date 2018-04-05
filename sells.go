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

Example:
 2bbf394c-193b-5b2a-9155-3b4732659ede

*/
type SellId string

// ListSells requires an account ID and returns an APIWalletTransferList struct
func (a *APIClient) ListSells(id AccountId) (sells APIWalletTransferList, err error) {
  path := pathHelper("/v2/accounts/%s/sells", id)
  err = a.Fetch("GET", path, nil, &sells)
  if err != nil {
    return
  }
  return
}

// ShowSell requires an account ID, buy ID and returns an APIWalletTransfer struct
func (a *APIClient) ShowSell(id AccountId, sid SellId) (sells APIWalletTransfer, err error) {
  path := pathHelper("/v2/accounts/%s/sells/%s", id, sid)
  err = a.Fetch("GET", path, nil, &sells)
  if err != nil {
    return
  }
  return
}

// PlaceSellOrder requires an account ID, APIWalletTransferOrder and returns an APIWalletTransfer struct
func (a *APIClient) PlaceSellOrder(id AccountId, order APIWalletTransferOrder) (sells APIWalletTransfer, err error) {

  path := pathHelper("/v2/accounts/%s/sells", id)
  err = a.Fetch("POST", path, order, &sells)
  if err != nil {
    return
  }
  return
}

// CommitSell requires an account ID, sell ID and returns an APIWalletTransfer struct
func (a *APIClient) CommitSell(id AccountId, sid SellId) (sells APIWalletTransfer, err error) {
  path := pathHelper("/v2/accounts/%s/buys/%s/commit", id, sid)
  err = a.Fetch("POST", path, nil, &sells)
  if err != nil {
    return
  }
  return
}
