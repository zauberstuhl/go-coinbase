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
   "data": {
     "currency": "BTC",
     "rates": {
       "AED": "36.73",
       "AFN": "589.50",
       "ALL": "1258.82",
       "AMD": "4769.49",
       "ANG": "17.88",
       "AOA": "1102.76",
       "ARS": "90.37",
       "AUD": "12.93",
       "AWG": "17.93",
       "AZN": "10.48",
       "BAM": "17.38",
       ...
     }
   }
 }

*/
type APIExchangeRatesData struct {
  Currency string
  Rates interface{}
}
type APIExchangeRates struct {
  Data []APIExchangeRatesData
}
// GetExchangeRates requires the currency as parameter and returns an APIExchangeRates struct
func (a *APIClient) GetExchangeRates(currency string) (rates APIExchangeRates, err error) {
  err = a.Fetch("GET", "/v2/exchange-rates?currency=" + currency, nil, &rates)
  if err != nil {
    return
  }
  return
}
