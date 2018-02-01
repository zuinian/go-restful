package models

import (
)

var (
	SchoolList map[string]*School
)

func init() {
	SchoolList = make(map[string]*School)
	s := School{"xiada","siming","12345"}
	SchoolList["xiada"] = &s
}


type School struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

func SchoolPost(s School) {
	SchoolList[s.Name] = &s
}

func SchoolsGet() map[string]*School {
	return SchoolList
}