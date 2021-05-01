package entity

import "errors"

type Planet struct {
	ID            ID
	Name          string
	Climate       string
	Ground        string
	FilmsQuantity int
}

func NewPlanet(name, climate, ground string) (*Planet, error) {
	p := &Planet{
		ID:      NewID(),
		Name:    name,
		Climate: climate,
		Ground:  ground,
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Planet) Validate() error {
	if p.Climate == "" || p.Ground == "" {
		return errors.New("change to err invalid entity")
	}
	return nil
}
