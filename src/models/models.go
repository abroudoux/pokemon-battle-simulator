package models

type Name struct {
    English  string `json:"english"`
    Japanese string `json:"japanese"`
    Chinese  string `json:"chinese"`
    French   string `json:"french"`
}

type TypeRelation struct {
    TypeName string `json:"typeName"`
    TypeId int `json:"typeId"`
}

type Type struct {
    Id         int    `json:"id"`
    Name    Name `json:"name"`
    Weaknessess []TypeRelation `json:"weaknesses"`
    Resistances []TypeRelation `json:"resistances"`
    Immunities  []TypeRelation `json:"immunities"`
}

type Base struct {
    HP        int `json:"HP"`
    Attack    int `json:"Attack"`
    Defense   int `json:"Defense"`
    SpAttack  int `json:"Sp. Attack"`
    SpDefense int `json:"Sp. Defense"`
    Speed     int `json:"Speed"`
}

type Pokemon struct {
    Id   int    `json:"id"`
    Name Name   `json:"name"`
    Type []Type `json:"type"`
    Base Base   `json:"base"`
}