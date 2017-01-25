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
   "data": [
     {
       "id": "AED",
       "name": "United Arab Emirates Dirham",
       "min_size": "0.01000000"
     },
     {
       "id": "AFN",
       "name": "Afghan Afghani",
       "min_size": "0.01000000"
     },
     {
       "id": "ALL",
       "name": "Albanian Lek",
       "min_size": "0.01000000"
     },
     {
       "id": "AMD",
       "name": "Armenian Dram",
       "min_size": "0.01000000"
     },
     ...
   }
 }

*/
type APICurrenciesData struct {
  Id string
  Name string
  Min_size float64
}
type APICurrencies struct {
  Data []APICurrenciesData
}
// GetCurrencies will return an APICurrencies struct
func (a *APIClient) GetCurrencies() (curr APICurrencies, err error) {
  err = a.Fetch("GET", "/v2/currencies", nil, &curr)
  if err != nil {
    return
  }
  return
}
