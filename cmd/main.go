package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/api/handler"
	"github.com/victor-schumacher/planets-B2W/api/integration/starwars"
	"github.com/victor-schumacher/planets-B2W/database/mongo"
	"github.com/victor-schumacher/planets-B2W/database/mongo/repository"
	"net/http"
	"time"
)

func main() {
	err := run()
	if err != nil {
		return
	}
}

func run() error {
	db, err := mongo.NewConnection("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
	}
	httpClient := http.Client{Timeout: time.Second * 2}
	planetRepo := repository.NewPlanet(db)
	cache := starwars.NewCache(planetRepo, httpClient)
	if err := cache.StartCron(); err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	planetHandler := handler.NewHandler(planetRepo)
	planetHandler.Handle(e)
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
