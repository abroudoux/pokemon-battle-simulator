package models

type Pokemon struct {
	Id int `json:"id"`
    Name struct {
        English string `json:"english"`
        Japanese string `json:"japanese"`
        Chinese string `json:"chinese"`
        French string `json:"french"`
    } `json:"name"`
	Type []string `json:"type"`
    Base struct {
        HP int `json:"HP"`
        Attack int `json:"Attack"`
        Defense int `json:"Defense"`
        SpAttack int `json:"Sp.Attack"`
        SpDefense int `json:"Sp.Defense"`
        Speed int `json:"Speed"`
    } `json:"base"`
}