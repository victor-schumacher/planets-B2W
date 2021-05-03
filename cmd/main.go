package main

import (
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/api/handler"
	"github.com/victor-schumacher/planets-B2W/api/integration/starwars"
	"github.com/victor-schumacher/planets-B2W/database/mongo"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	db, err := mongo.NewConnection("mongodb://127.0.0.1:27017")
	if err != nil {
		return err
	}
	httpClient := http.Client{Timeout: time.Second * 2}
	planetRepo := repository.NewPlanet(db)

	cache := starwars.NewCache(planetRepo, httpClient)
	if err := cache.StartCron(); err != nil {
		return err
	}

	e := echo.New()
	planetHandler := handler.NewPlanet(planetRepo)
	planetHandler.Handle(e)
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}
