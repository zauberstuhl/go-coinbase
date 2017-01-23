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
 * {
 *   "data": {
 *     "iso": "2015-06-23T18:02:51Z",
 *     "epoch": 1435082571
 *   }
 * }
 */
type APITimeData struct {
  Iso string
  Epoch int64
}
type APITime struct {
  Data APITimeData
}
// Get current time
func (a *APIClient) GetCurrentTime() (time APITime, err error) {
  err = a.Fetch("GET", "/v2/time", nil, &time)
  if err != nil {
    return
  }
  return
}
