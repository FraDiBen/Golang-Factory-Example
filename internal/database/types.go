package database

import "strings"

type DBs int

const (
	// Undefined : since iota starts with 0, the first value
	// defined here will be the default
	Undefined DBs = iota
	Mysql
	Mongo
)

func (d DBs) String() string {
	switch d {
	default:
		return "undefined"
	case Mysql:
		return "mysql"
	case Mongo:
		return "mongo"
	}
}

func DBFrom(input string) DBs {
	switch strings.ToLower(input) {
	case "mysql":
		return Mysql
	case "mongo":
		return Mongo
	default:
		return Undefined
	}
}
