package models

import (
)

var (
	SchoolList map[string]*School
)

func init() {
	SchoolList = make(map[string]*School)
	s := School{"xiada","siming","12345"}
	SchoolList["0"] = &s
}


type School struct {
	Name string
	Address string
	Phone string
}

func AddSchool(s School) string {
	return s.Name
}

func GetAllSchools() map[string]*School {
	return SchoolList
}