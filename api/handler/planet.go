package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
	"github.com/victor-schumacher/planets-B2W/entity"
	"net/http"
)

type Manager struct {
	planetRepo repository.Planet
}

func NewPlanet(planet repository.Planet) Manager {
	return Manager{planetRepo: planet}
}

func (m Manager) listAll(c echo.Context) error {
	planets, err := m.planetRepo.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, planets)
}

func (m Manager) findOne(c echo.Context) error {
	searchCriteria := c.Param("searchCriteria")
	fmt.Println(searchCriteria)
	if !isSearchCriteriaAllowed(searchCriteria) {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Errorf("search key %q not allowed", searchCriteria),
		)
	}

	search := c.Param("search")
	p, err := m.planetRepo.FindOne(searchCriteria, search)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

func (m Manager) add(c echo.Context) error {
	p := entity.Planet{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	planet, err := entity.NewPlanet(p.Name, p.Climate, p.Terrain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = m.planetRepo.Save(planet); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (m Manager) deletePlanet(c echo.Context) error {
	id := c.Param("id")
	if err := m.planetRepo.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func isSearchCriteriaAllowed(searchCriteria string) bool {
	if searchCriteria != "id" &&
		searchCriteria != "name" {
		return false
	}
	return true
}

func (m Manager) Handle(e *echo.Echo) {
	p := e.Group("/planets")
	p.GET("", m.listAll)
	p.GET("/:searchCriteria/:search", m.findOne)
	p.POST("", m.add)
	p.DELETE("/:id", m.deletePlanet)
}
