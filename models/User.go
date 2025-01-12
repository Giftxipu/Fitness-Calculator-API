package models

type User struct {
  Weight float64 `json:"weight"`;
  Height float64 `json:"height"`;
  Age int `json:"age"`;
  Gender string `json:"gender"`;
}
