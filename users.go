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
   "id": "9da7a204-544e-5fd1-9a12-61176c5d4cd8",
   "name": "User One",
   "username": "user1",
   "profile_location": null,
   "profile_bio": null,
   "profile_url": "https://coinbase.com/user1",
   "avatar_url": "https://images.coinbase.com/avatar?h=vR%2FY8igBoPwuwGren5JMwvDNGpURAY%2F0nRIOgH%2FY2Qh%2BQ6nomR3qusA%2Bh6o2%0Af9rH&s=128",
   "resource": "user",
   "resource_path": "/v2/user"
 }

wallet:user:read permission

 {
   ...
   "time_zone": "Pacific Time (US & Canada)",
   "native_currency": "USD",
   "bitcoin_unit": "bits",
   "country": {
     "code": "US",
     "name": "United States"
   },
   "created_at": "2015-01-31T20:49:02Z"
 }

wallet:user:email permission

 {
   ...
   "email": "user1@example.com"
 }

*/
type APIUserCountry struct {
  Code string
  Name string
}
type APIUserData struct {
  Id string
  Name string
  Username string
  Profile_location string
  Profile_bio string
  Profile_url string
  Avatar_url string
  Resource string
  Resource_path string
  // wallet:user:read permission
  Time_zone string
  Native_currency string
  Bitcoin_unit string
  Country APIUserCountry
  Created_at string
  // wallet:user:email permission
  Email string
}
type APIUser struct {
  Data APIUserData
  Errors []Error
}
// GetUser requires a user ID and returns an APIUser struct
func (a *APIClient) GetUser(id string) (user APIUser, err error) {
  err = a.Fetch("GET", "/v2/users/" + id, nil, &user)
  if err != nil {
    return
  }
  return
}

// GetCurrentUser returns an APIUser struct
func (a *APIClient) GetCurrentUser() (user APIUser, err error) {
  err = a.Fetch("GET", "/v2/user", nil, &user)
  if err != nil {
    return
  }
  return
}

/*

Example Response:

 {
   "data": {
     "method": "oauth",
     "scopes": [
         "wallet:user:read",
         "wallet:user:email"
     ],
     "oauth_meta": {}
   }
 }

*/
type APIUserAuthData struct {
  Method string
  Scopes []string
  Oauth_meta interface{}
}
type APIUserAuth struct {
  Data APIUserAuthData
}
// GetCurrentUserAuth returns an APIUserAuth struct
func (a *APIClient) GetCurrentUserAuth() (auth APIUserAuth, err error) {
  err = a.Fetch("GET", "/v2/user/auth", nil, &auth)
  if err != nil {
    return
  }
  return
}

// UpdateCurrentUser requires a new username and returns an APIUser struct
func (a *APIClient) UpdateCurrentUser(name string) (user APIUser, err error) {
  var body APIName
  body.Name = name
  err = a.Fetch("PUT", "/v2/user", body, &user)
  if err != nil {
    return
  }
  return
}
