package main

import "gopkg.in/mgo.v2"

type DB struct {
	colDrugs *mgo.Collection
	session  *mgo.Session
}

func CreateDB() *DB {
	db := DB{}
	session, err := mgo.Dial("5.200.55.232")
	db.session = session
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Eventual, true)
	db.colDrugs = session.DB("medproj").C("drugs")
	return &db
}
