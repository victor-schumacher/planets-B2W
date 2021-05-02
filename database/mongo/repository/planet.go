package repository

import (
	"github.com/victor-schumacher/planets-B2W/api/integration/starwars"
	"github.com/victor-schumacher/planets-B2W/database/mongo"
	"github.com/victor-schumacher/planets-B2W/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Planet interface {
	Save(planet entity.Planet) error
	SaveCache(planet starwars.PlanetCache) error
	FindAll() ([]entity.Planet, error)
	FindOne(searchCriteria string, search interface{}) (entity.Planet, error)
	filmsQuantity(planetName string) (int, error)
	Delete(ID string) error
}

type PlanetRepo struct {
	session *mgo.Session
}

func NewPlanet(dbSession *mgo.Session) PlanetRepo {
	return PlanetRepo{session: dbSession}
}

func (db PlanetRepo) Save(planet entity.Planet) error {
	s := db.getFreshSession()
	defer s.Close()


	q, err  := db.filmsQuantity(planet.Name)
	if err != nil {
		return err
	}
	planet.FilmsQuantity = q
	return s.DB(mongo.DB).C(mongo.PLANETS).Insert(planet)
}

func (db PlanetRepo) FindAll() ([]entity.Planet, error) {
	s := db.getFreshSession()
	defer s.Close()

	var planets []entity.Planet
	if err := s.DB(mongo.DB).C(mongo.PLANETS).Find(nil).All(&planets); err != nil {
		return nil, err
	}
	return planets, nil
}

func (db *PlanetRepo) getFreshSession() *mgo.Session {
	return db.session.Copy()
}

func (db PlanetRepo) FindOne(searchCriteria string, search interface{}) (entity.Planet, error) {
	s := db.getFreshSession()
	defer s.Close()
	p := entity.Planet{}
	if err := s.DB(mongo.DB).C(mongo.PLANETS).
		Find(bson.M{searchCriteria: search}).
		One(&p);
		err != nil {
		return p, err
	}
	return p, nil
}

func (db PlanetRepo) Delete(ID string) error {
	s := db.getFreshSession()
	defer s.Close()
	if err := s.DB(mongo.DB).C(mongo.PLANETS).Remove(bson.M{"id": ID}); err != nil {
		return err
	}

	return nil
}

func (db PlanetRepo) SaveCache(planet starwars.PlanetCache) error {
	s := db.getFreshSession()
	defer s.Close()
	return s.DB(mongo.DB).C(mongo.PLANETSCACHE).Insert(planet)
}

func (db PlanetRepo) filmsQuantity(planetName string) (int, error) {
	s := db.getFreshSession()
	defer s.Close()
	p := starwars.PlanetCache{}
	if err := s.DB(mongo.DB).C(mongo.PLANETSCACHE).
		Find(bson.M{"name":planetName}).
		One(&p);
		err != nil {
		return p.FilmsQuantity, err
	}
	return p.FilmsQuantity, nil
}
