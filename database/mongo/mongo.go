package mongo

import "gopkg.in/mgo.v2"

const (
	DB      = "myplanets"
	PLANETS = "planets"
	PLANETSCACHE = "planetsCache"
)

func NewConnection(connection string) (*mgo.Session, error) {
	c, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}
	return c, nil
}
