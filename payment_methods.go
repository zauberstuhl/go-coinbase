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
   ...
   "limits": {
     "buy": [
       {
         "period_in_days": 1,
         "total": {
           "amount": "3000.00",
           "currency": "USD"
         },
         "remaining": {
           "amount": "3000.00",
           "currency": "USD"
         }
       }
     ],
     "instant_buy": [
       {
         "period_in_days": 7,
         "total": {
           "amount": "0.00",
           "currency": "USD"
         },
         "remaining": {
           "amount": "0.00",
           "currency": "USD"
         }
       }
     ],
     "sell": [
       {
         "period_in_days": 1,
         "total": {
           "amount": "3000.00",
           "currency": "USD"
         },
         "remaining": {
           "amount": "3000.00",
           "currency": "USD"
         }
       }
     ],
     "deposit": [
       {
         "period_in_days": 1,
         "total": {
           "amount": "3000.00",
           "currency": "USD"
         },
         "remaining": {
           "amount": "3000.00",
           "currency": "USD"
         }
       }
     ]
   },
 }

*/
type APIPaymentMethodsLimitsData struct {
  Period_in_days int
  Total APIBalance
  Remaining APIBalance
}
type APIPaymentMethodsLimits struct {
  Buy []APIPaymentMethodsLimitsData
  Instant_buy []APIPaymentMethodsLimitsData
  Sell []APIPaymentMethodsLimitsData
  Deposit []APIPaymentMethodsLimitsData
}

/*

Example Response:

 {
   "id": "83562370-3e5c-51db-87da-752af5ab9559",
   "type": "ach_bank_account",
   "name": "International Bank *****1111",
   "currency": "USD",
   "primary_buy": true,
   "primary_sell": true,
   "allow_buy": true,
   "allow_sell": true,
   "allow_deposit": true,
   "allow_withdraw": true,
   "instant_buy": false,
   "instant_sell": false,
   "created_at": "2015-01-31T20:49:02Z",
   "updated_at": "2015-02-11T16:53:57-08:00",
   "resource": "payment_method",
   "resource_path": "/v2/payment-methods/83562370-3e5c-51db-87da-752af5ab9559"
 }

*/
type APIPaymentMethodsData struct {
  ID string
  Type string
  Name string
  Currency string
  Primary_buy bool
  Primary_sell bool
  Allow_buy bool
  Allow_sell bool
  Allow_deposit bool
  Allow_withdraw bool
  Instant_buy bool
  Instant_sell bool
  Created_at string
  Updated_at string
  Resource string
  Resource_path string
  Limits APIPaymentMethodsLimits
}
type APIPaymentMethods struct {
  Pagination APIPagination
  Data []APIPaymentMethodsData
  Errors []Error
}
type APIPaymentMethod struct {
  Data APIPaymentMethodsData
  Errors []Error
}

// ListPaymentMethods returns an APIPaymentMethods struct
func (a *APIClient) ListPaymentMethods() (methods APIPaymentMethods, err error) {
  err = a.Fetch("GET", "/v2/payment-methods", nil, &methods)
  if err != nil {
    return
  }
  return
}

/*

Example:
 83562370-3e5c-51db-87da-752af5ab9559

*/
type PaymentMethodId string

// ListPaymentMethods returns an APIPaymentMethods struct
func (a *APIClient) ShowPaymentMethod(id PaymentMethodId) (method APIPaymentMethod, err error) {
  path := pathHelper("/v2/payment-methods/%s", id)
  err = a.Fetch("GET", path, nil, &method)
  if err != nil {
    return
  }
  return
}
