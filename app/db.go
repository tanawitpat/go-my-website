package app

import (
	"time"

	"github.com/globalsign/mgo"
)

func GetMongoSession() (*mgo.Session, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  5 * time.Second,
		Database: "admin",
		Username: "admin",
		Password: "passw0rd",
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	return session, err
}
