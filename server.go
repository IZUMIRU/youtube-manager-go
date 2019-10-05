package main

import (
	"youtube-manager-go/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New() // 1
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080")) // 2
}
