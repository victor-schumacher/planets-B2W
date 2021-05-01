package entity

import "errors"

type Planet struct {
	ID            ID     `json:"id"`
	Name          string `json:"name" validate:"required"`
	Climate       string `json:"climate" validate:"required"`
	Ground        string `json:"ground" validate:"required"`
	FilmsQuantity int    `json:"filmsQuantity" validate:"required"`
}

func NewPlanet(name, climate, ground string) (Planet, error) {
	p := Planet{
		ID:      NewID(),
		Name:    name,
		Climate: climate,
		Ground:  ground,
	}
	if err := p.Validate(); err != nil {
		return p, err
	}
	return p, nil
}

func (p *Planet) Validate() error {
	if p.Climate == "" || p.Ground == "" {
		return errors.New("change to err invalid entity")
	}
	return nil
}
