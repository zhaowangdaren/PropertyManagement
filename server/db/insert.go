package db

import (
	"./table"
	mgo "gopkg.in/mgo.v2"
)

//InsertStreet insert
func InsertStreet(db *mgo.Database, street table.Street) interface{} {
	return table.InsertStreetTB(db, street)
}
