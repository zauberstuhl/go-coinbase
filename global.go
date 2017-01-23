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
 *   "pagination": {
 *     "ending_before": null,
 *     "starting_after": null,
 *     "limit": 25,
 *     "order": "desc",
 *     "previous_uri": null,
 *     "next_uri": null
 *   },
 */
type APIPagination struct {
  Ending_before string
  Starting_after string
  Limit int
  Order string
  Previous_uri string
  Next_uri string
}

/*
 *   "balance": {
 *       "amount": "39.59000000",
 *       "currency": "BTC"
 *   },
 */
type APIBalance struct {
  Amount float64 `json:",string"`
  Currency string
}

