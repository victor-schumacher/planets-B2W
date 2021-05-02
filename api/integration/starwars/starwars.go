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
}

type PlanetCache struct {
	Name          string `json:"name"`
	FilmsQuantity int    `json:"filmsquantity"`
}

type PlanetsCache []PlanetCache

func Planets() ([]PlanetCache, error) {
	next := "https://swapi.dev/api/planets/"
	var psc []PlanetCache
	for {
		r, err := swClient.Get(next)
		if err != nil {
			return nil, err
		}
		p := Response{}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			return nil, err
		}

		pc := PlanetCache{}

		for _, planet := range p.Planets {
			pc.FilmsQuantity = len(planet.Films)
			pc.Name = planet.Name
			psc = append(psc, pc)
		}

		if p.Next == "" {
			break
		}
		next = p.Next
	}


	return psc, nil
}
