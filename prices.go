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

// TODO wrong naming convention
type APIBalanceData struct {
  Data APIBalance
}
// Get buy price
func (a *APIClient) GetBuyPrice(from, to string) (balance APIBalanceData, err error) {
  err = a.Fetch("GET", "/v2/prices/" + from + "-" + to + "/buy", nil, &balance)
  if err != nil {
    return
  }
  return
}

// Get sell price
func (a *APIClient) GetSellPrice(from, to string) (balance APIBalanceData, err error) {
  err = a.Fetch("GET", "/v2/prices/" + from + "-" + to + "/sell", nil, &balance)
  if err != nil {
    return
  }
  return
}

// Get spot price
func (a *APIClient) GetSpotPrice(from, to string) (balance APIBalanceData, err error) {
  err = a.Fetch("GET", "/v2/prices/" + from + "-" + to + "/spot", nil, &balance)
  if err != nil {
    return
  }
  return
}
