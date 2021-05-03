package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/victor-schumacher/planets-B2W/internal/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePlanet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/planets", strings.NewReader(mock.PlanetJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewPlanet(mock.PlanetRepoMock{})

	if assert.NoError(t, h.add(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestDeletePlanet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/planets/d647faa7-c945-443f-9fc9-a06e7196e8fc", strings.NewReader(mock.PlanetJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewPlanet(mock.PlanetRepoMock{})

	if assert.NoError(t, h.deletePlanet(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestFindOne(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/planets/id/d647faa7-c945-443f-9fc9-a06e7196e8fc", strings.NewReader(mock.PlanetJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("searchCriteria")
	c.SetParamValues("id")
	h := NewPlanet(mock.PlanetRepoMock{})

	if assert.NoError(t, h.findOne(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestFindAll(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/planets", strings.NewReader(mock.PlanetJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewPlanet(mock.PlanetRepoMock{})

	if assert.NoError(t, h.listAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
