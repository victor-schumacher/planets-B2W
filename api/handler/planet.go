package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
	"github.com/victor-schumacher/planets-B2W/entity"
	"net/http"
)

type Manager struct {
	planetRepo repository.Planet
}

func NewHandler(planet repository.Planet) Manager {
	return Manager{planetRepo: planet}
}

func (m Manager) listPlanets(c echo.Context) error {
	planets, err := m.planetRepo.FindAll()
	if err != nil {
		return err
	}

	if err := c.JSON(http.StatusOK, planets); err != nil {
		return err
	}
	return nil
}

func (m Manager) findPlanet(c echo.Context) error {
	searchCriteria := c.Param("searchCriteria")
	if !isSearchCriteriaAllowed(searchCriteria) {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New(
			"search key not allowed").Error(),
		)
	}

	search := c.Param("search")
	p, err := m.planetRepo.FindOne(searchCriteria, search)
	if err != nil {
		return echo.NewHTTPError(http.StatusAlreadyReported, err.Error())
	}
	if err := c.JSON(http.StatusOK, p); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return nil
}

func (m Manager) addPlanet(c echo.Context) error {
	p := entity.Planet{}
	err := c.Bind(&p)
	if err != nil {
		return err
	}
	planet, err := entity.NewPlanet(p.Name, p.Climate, p.Ground)
	if err != nil {
		return err
	}
	if err = m.planetRepo.Save(planet); err != nil {
		return err
	}
	return nil
}

func (m Manager) deletePlanet(c echo.Context) error {
	id := c.Param("id")
	if err := m.planetRepo.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusAlreadyReported, err.Error())
	}
	return nil
}

func isSearchCriteriaAllowed(searchCriteria string) bool {
	if searchCriteria != "id" &&
		searchCriteria != "filmsQuantity" {
		return false
	}
	return true
}

func (m Manager) Handle(e *echo.Echo) {
	p := e.Group("/planets")
	p.GET("", m.listPlanets)
	p.GET("/:searchCriteria/:search", m.findPlanet)
	p.POST("", m.addPlanet)
	p.DELETE("/:id", m.deletePlanet)
}
