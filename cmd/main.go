package main

import (
	"github.com/labstack/echo/v4"
	"github.com/victor-schumacher/planets-B2W/api/handler"
)

func main(){
	e := echo.New()
	handler.MakePlanetHandlers(e)
	e.Logger.Fatal(e.Start(":1323"))
}