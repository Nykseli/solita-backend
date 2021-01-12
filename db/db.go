package db

import "io/ioutil"

// DB contains the data from names.json as a raw string
type DB struct {
	data string
}

// Singleton instance
var dbInstance *DB

// GetDBInstance returns the singleton DB instance
func GetDBInstance() *DB {
	if dbInstance == nil {
		dbInstance = &DB{}
		data, err := ioutil.ReadFile("names.json")
		if err != nil {
			panic("Couldn't load names.json")
		}
		dbInstance.data = string(data)
	}

	return dbInstance
}
