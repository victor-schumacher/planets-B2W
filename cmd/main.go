package main

import (
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/api/handler"
	"github.com/victor-schumacher/planets-B2W/api/integration/starwars"
	"github.com/victor-schumacher/planets-B2W/database/mongo"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
)

func main(){
	starwars.Planets()
	db, _ := mongo.NewConnection("")
	planetRepo := repository.NewPlanet(db)
	e := echo.New()
	planetHandler := handler.NewHandler(planetRepo)
	planetHandler.Handle(e)
	e.Logger.Fatal(e.Start(":1323"))
}