package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/api/handler"
	"github.com/victor-schumacher/planets-B2W/api/integration/starwars"
	"github.com/victor-schumacher/planets-B2W/database/mongo"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
)

func main(){
	planets, err := starwars.Planets()
	if err != nil {
		fmt.Println(err)
	}

	db, err := mongo.NewConnection("")
	if err != nil {
		fmt.Println(err)
	}
	planetRepo := repository.NewPlanet(db)

	for _, p := range planets {
		err = planetRepo.SaveCache(p)
		if err != nil {
			fmt.Println(err)
		}
	}
	e := echo.New()
	planetHandler := handler.NewHandler(planetRepo)
	planetHandler.Handle(e)
	e.Logger.Fatal(e.Start(":1323"))
}