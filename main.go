package main

import (
	. "factoryExampleGo/internal"
	db "factoryExampleGo/internal/database"
	"os"
	"strings"
)

//main is the Composition Root
func main() {
	// create a factory of DBs based on some configs, in this case ENV variables
	var factory db.DBFactory
	dbType := db.DBFrom(getEnv("DB", db.Mongo.String()))
	if dbType == db.Mongo {
		factory = db.MongoFactory{}
	} else {
		factory = db.MySQLFactory{}
	}
	// run is the real main, can be tested
	Run(factory)
}

func getEnv(env string, def string) string {
	e := strings.TrimSpace(os.Getenv(env))
	if len(e) <= 0 {
		return def
	}
	return e
}
