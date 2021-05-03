package entity

type Planet struct {
	ID            ID     `json:"id"`
	Name          string `json:"name" validate:"required"`
	Climate       string `json:"climate" validate:"required"`
	Terrain       string `json:"terrain" validate:"required"`
	FilmsQuantity int    `json:"filmsQuantity" validate:"required"`
}

type PlanetCache struct {
	Name          string `json:"name"`
	FilmsQuantity int    `json:"filmsquantity"`
}

func NewPlanet(name, climate, terrain string) (Planet, error) {
	p := Planet{
		ID:      NewID(),
		Name:    name,
		Climate: climate,
		Terrain: terrain,
	}
	if err := p.Validate(); err != nil {
		return p, err
	}
	return p, nil
}

func (p *Planet) Validate() error {
	if p.Climate == "" || p.Terrain == "" || p.Name == "" {
		return ErrInvalidEntity
	}
	return nil
}
