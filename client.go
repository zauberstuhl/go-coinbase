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
  "fmt"
  "encoding/json"
  "net/http"
  "crypto/sha256"
  "crypto/hmac"
  "strconv"
  "io"
)

const (
  ENDPOINT = "https://api.coinbase.com"
  API_VERSION = "2016-03-08"
  API_TIME = "/v2/time"
)

type APIClient struct {
  Key string
  Secret string
}

type APIClientEpoch struct {
  Data struct {
    Epoch int64
  }
}

func (a *APIClient) Fetch(method, path string, body io.Reader, result interface{}) error {
  client := &http.Client{}
  req, err := http.NewRequest(method, ENDPOINT + path, body)
  if err != nil {
    return err
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("CB-VERSION", API_VERSION)
  // do not authenticate on public time api call
  if path[len(path)-4:] != "time" {
    err = a.Authenticate(path, req)
    if err != nil {
      return err
    }
  }

  resp, err := client.Do(req)
  if err != nil {
    return err
  }
  err = json.NewDecoder(resp.Body).Decode(result)
  if err != nil {
    return err
  }
  return nil
}

func (a *APIClient) Authenticate(path string, req *http.Request) error {
  time, err := a.GetCurrentTime()
  if err != nil {
    return err
  }
  timestamp := strconv.FormatInt(time.Data.Epoch, 10)
  message := timestamp + req.Method + path

  sha := sha256.New
  h := hmac.New(sha, []byte(a.Secret))
  h.Write([]byte(message))

  signature := fmt.Sprintf("%x", h.Sum(nil))

  req.Header.Set("CB-ACCESS-KEY", a.Key)
  req.Header.Set("CB-ACCESS-SIGN", signature)
  req.Header.Set("CB-ACCESS-TIMESTAMP", timestamp)

  return nil
}
