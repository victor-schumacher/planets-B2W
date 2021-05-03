package starwars

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
	"github.com/victor-schumacher/planets-B2W/entity"
	"net/http"
	"time"
)

type Response struct {
	Next    string   `json:"next"`
	Planets []Planet `json:"results"`
}

type Planet struct {
	Name    string   `json:"name"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"films"`
}

type PlanetsCache []entity.PlanetCache

type Manager struct {
	planetRepo repository.Planet
	client     http.Client
}

func NewCache(
	planet repository.Planet,
	client http.Client,
) Manager {
	return Manager{
		planetRepo: planet,
		client:     client,
	}
}

func (m Manager) StartCron() error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Second().Do(m.cacheUpdate)
	if err != nil {
		return err
	}
	s.StartAsync()
	return nil
}

func (m Manager) cacheUpdate() error {
	planets, err := m.planets()
	if err != nil {
		return err
	}
	for _, p := range planets {
		err = m.planetRepo.SaveCache(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Manager) planets() ([]entity.PlanetCache, error) {
	next := "https://swapi.dev/api/planets/"
	var psc []entity.PlanetCache
	for {
		r, err := m.client.Get(next)
		if err != nil {
			return nil, err
		}
		p := Response{}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			return nil, err
		}

		pc := entity.PlanetCache{}
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
