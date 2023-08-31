package models

type User struct {
	Id   int
	Name string `json:"name"`
	Last string `json:"last"`
}
