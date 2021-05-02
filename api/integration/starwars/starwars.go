package starwars

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	URL = "https://swapi.dev/api/planets/"
)

var swClient = http.Client{Timeout: time.Second * 2}

type Response struct {
	Next    string   `json:"next"`
	Planets []Planet `json:"results"`
}

type Planet struct {
	Name          string   `json:"name"`
	Climate       string   `json:"climate"`
	Terrain       string   `json:"terrain"`
	Films         []string `json:"films"`
	FilmsQuantity int
}

type PlanetCache struct {
	Name          string `json:"name"`
	FilmsQuantity int    `json:"filmsQuantity"`
}

func Planets() ([]PlanetCache, error) {
	r, err := swClient.Get(URL)
	if err != nil {
		return nil, err
	}
	p := Response{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		return nil, err
	}

	pc := PlanetCache{}
	var psc []PlanetCache
	for _, planet := range p.Planets {
		pc.FilmsQuantity = len(planet.Films)
		pc.Name = planet.Name
		psc = append(psc, pc)
	}
	return psc, nil
}
