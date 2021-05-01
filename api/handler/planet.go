package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
	"github.com/victor-schumacher/planets-B2W/entity"
)

type Planet struct {
	Name    string `json:"name" validate:"required"`
	Climate string `json:"climate" validate:"required"`
	Ground  string `json:"ground" validate:"required"`
}

type Manager struct {
	planetRepo repository.Planet
}

func NewHandler(planet repository.Planet) Manager {
	return Manager{planetRepo: planet}
}

func (m Manager) listPlanets(c echo.Context) error {
	return nil
}

func (m Manager) findPlanet(c echo.Context) error {
	searchCriteria := c.Param("searchCriteria")
	search := c.Param("search")
	switch searchCriteria {
	case "id":
		fmt.Println("find by id" + search)
	case "name":
		fmt.Println("find by id" + search)
	default:
		fmt.Println("find by name" + search)
	}
	return nil
}

func (m Manager) addPlanet(c echo.Context) error {
	p := Planet{}
	err := c.Bind(&p)
	if err != nil {
		return err
	}

	pe := entity.Planet{
		Name:          p.Name,
		Climate:       p.Climate,
		Ground:        p.Ground,
		FilmsQuantity: 0,
	}
	_, err = m.planetRepo.Save(pe)
	if err != nil {
		fmt.Print(err)
	}
	return nil
}

func (m Manager) deletePlanet(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return nil
}

func (m Manager) Handle(e *echo.Echo) {
	p := e.Group("/planets")
	p.GET("", m.listPlanets)
	p.GET("/:searchCriteria/:search", m.findPlanet)
	p.POST("", m.addPlanet)
	p.DELETE("/:id", m.deletePlanet)
	fmt.Println("here")
}
