package coinbase

import (
  "testing"
)

func TestAddresses(t *testing.T) {
  s := serverHelper(`
 {
   "data": {
       "id": "dd3183eb-af1d-5f5d-a90d-cbff946435ff",
       "address": "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa",
       "name": "One off payment",
       "created_at": "2015-01-31T20:49:02Z",
       "updated_at": "2015-03-31T17:25:29-07:00",
       "network": "bitcoin",
       "resource": "address",
       "resource_path": "/v2/accounts/2bbf394c-193b-5b2a-9155-3b4732659ede/addresses/dd3183eb-af1d-5f5d-a90d-cbff946435ff"
   }
 }
  `)
  defer s.Close()

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: s.URL,
  }

  address, err := a.GetAddress(
    "dd3183eb-af1d-5f5d-a90d-cbff946435ff",
    "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa",
  )
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  if address.Data.Id != "dd3183eb-af1d-5f5d-a90d-cbff946435ff" {
    t.Error("Expected dd3183eb-af1d-5f5d-a90d-cbff946435ff, got ",
      address.Data.Id)
  }
  if address.Data.Address != "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa" {
    t.Error("Expected mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa, got ",
      address.Data.Address)
  }
  if address.Data.Created_at.Year() != 2015 {
    t.Error("Expected 2015, got ", address.Data.Created_at.Year())
  }
  if address.Data.Updated_at.Year() != 2015 {
    t.Error("Expected 2015, got ", address.Data.Created_at.Year())
  }
}
