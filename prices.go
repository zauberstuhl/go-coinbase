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
  Errors []Error
}
// GetBuyPrice requires a ConfigPrice struct as parameter
// and returns an APIBalanceData struct
func (a *APIClient) GetBuyPrice(c ConfigPrice) (balance APIBalanceData, err error) {
  err = a.Fetch("GET", "/v2/prices/" + c.From +
    "-" + c.To + "/buy", nil, &balance)
  if err != nil {
    return
  }
  return
}

// GetSellPrice requires a ConfigPrice struct as parameter
// and returns an APIBalanceData struct
func (a *APIClient) GetSellPrice(c ConfigPrice) (balance APIBalanceData, err error) {
  err = a.Fetch("GET", "/v2/prices/" + c.From +
    "-" + c.To + "/sell", nil, &balance)
  if err != nil {
    return
  }
  return
}

// GetSpotPrice requires a ConfigPrice struct as parameter
// and returns an APIBalanceData struct. If you define a date
// you will recieve the spot price for the given date.
func (a *APIClient) GetSpotPrice(c ConfigPrice) (balance APIBalanceData, err error) {
  var date string = ""
  if c.Date.Year() != 1 {
    date = "?date=" + c.Date.Format("2006-01-02")
  }
  err = a.Fetch("GET", "/v2/prices/" + c.From +
    "-" + c.To + "/spot" + date, nil, &balance)
  if err != nil {
    return
  }
  return
}
