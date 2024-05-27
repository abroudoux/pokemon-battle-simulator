package models

type Pokemon struct {
	Id int `json:"id"`
	Name string  `json:"name"`
	Type []string `json:"type"`
	Base []int `json:"base"`
}