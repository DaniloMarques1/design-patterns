package main

import "errors"

const (
	MONGO = iota
	SQL
)

type Person struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type Database interface {
	Save(*Person) error
	FindByName(string) (*Person, error)
}

func GetDatabase(dbType uint) (Database, error) {
	var database Database
	var err error
	switch dbType {
	case MONGO:
		database, err = NewMongoDatabase()
	case SQL:
		database, err = NewSqlDatabase()
	default:
		err = errors.New("Invalid db type")
	}

	if err != nil {
		return nil, err
	}

	return database, nil
}
