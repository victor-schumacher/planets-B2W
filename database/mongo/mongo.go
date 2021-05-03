package mongo

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

const (
	DB           = "myplanets"
	PLANETS      = "planets"
	PLANETSCACHE = "planetsCache"
)

var DBURL = os.Getenv("MONGO_URL")

func NewConnection(connection string) (*mgo.Session, error) {
	c, err := mgo.Dial(connection)
	if err != nil {
		log.Fatalln("cannot connect to mongo database:" + err.Error())
	}
	return c, nil
}
