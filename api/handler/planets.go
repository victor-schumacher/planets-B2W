package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func listPlanets(c echo.Context) error {
	return nil
}

func findPlanet(c echo.Context) error {
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

func addPlanet(c echo.Context) error {
	fmt.Println("add")
	return nil
}

func deletePlanet(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	return nil
}

func MakePlanetHandlers(e *echo.Echo) {
	p := e.Group("/planets")
	p.GET("/", listPlanets)
	p.GET("/:searchCriteria/:search", findPlanet)
	p.POST("/", addPlanet)
	p.DELETE("/:id", deletePlanet)
}
