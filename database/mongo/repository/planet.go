package repository

import (
	"github.com/victor-schumacher/planets-B2W/database/mongo"
	"github.com/victor-schumacher/planets-B2W/entity"
	"gopkg.in/mgo.v2"
)

type Planet interface {
	Save(planet entity.Planet) (entity.ID, error)
	// FindAll() ([]entity.Planet, error)
	// FindByName(ID string) (entity.Planet, error)
	// FindByID(name string) (entity.Planet, error)
	// Delete(ID string) error
}

type PlanetRepo struct {
	session *mgo.Session
}

func NewPlanet(dbSession *mgo.Session) PlanetRepo {
	return PlanetRepo{session: dbSession}
}

func (db PlanetRepo) Save(planet entity.Planet) (entity.ID, error) {
	s := db.getFreshSession()
	defer s.Close()
	return entity.NewID(), s.DB(mongo.DB).C(mongo.PLANETS).Insert(planet)
}

func (db *PlanetRepo) getFreshSession() *mgo.Session {
	return db.session.Copy()
}
