package mock

import "github.com/victor-schumacher/planets-B2W/entity"

type PlanetRepoMock struct {
}

func (PlanetRepoMock) Save(planet entity.Planet) error {
	return nil
}
func (PlanetRepoMock) SaveCache(planet entity.PlanetCache) error {
	return nil
}
func (PlanetRepoMock) FindAll() ([]entity.Planet, error) {
	var ps []entity.Planet
	return ps, nil
}
func (PlanetRepoMock) FindOne(searchCriteria string, search interface{}) (entity.Planet, error) {
	p := entity.Planet{}
	return p, nil
}
func (PlanetRepoMock) Delete(ID string) error {
	return nil
}

const PlanetJson = `{
    "name":"Tatooine",
    "climate":"arid",
    "terrain":"desert"
}`
