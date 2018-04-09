package coinbase

import (
  "testing"
)

func TestAddress(t *testing.T) {
  method := "GET"
  uri := "/v2/accounts/dd3183eb-af1d-5f5d-a90d-cbff946435ff/addresses/mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa"
  body := `{
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
  }`
  s := serverHelper(t, &method, &uri, &body)
  defer s.Close()

  a := APIClient{
    Key: "123",
    Secret: "123456",
    Endpoint: s.URL,
  }

  add, err := a.GetAddress(
    "dd3183eb-af1d-5f5d-a90d-cbff946435ff",
    "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  address(t, add.Data)

  method = "POST"
  uri = "/v2/accounts/mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa/addresses"
  add, err = a.CreateAddress(
    "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa",
    "One off payment")
  if err != nil {
    t.Error("Expected nil, got ", err.Error())
  }
  address(t, add.Data)
}

func address(t *testing.T, address APIAddressData) {
  if address.Id != "dd3183eb-af1d-5f5d-a90d-cbff946435ff" {
    t.Error("Expected dd3183eb-af1d-5f5d-a90d-cbff946435ff, got ", address.Id)
  }
  if address.Address != "mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa" {
    t.Error("Expected mswUGcPHp1YnkLCgF1TtoryqSc5E9Q8xFa, got ", address.Address)
  }
  if address.Created_at.Year() != 2015 {
    t.Error("Expected 2015, got ", address.Created_at.Year())
  }
  if address.Updated_at.Year() != 2015 {
    t.Error("Expected 2015, got ", address.Created_at.Year())
  }
}
