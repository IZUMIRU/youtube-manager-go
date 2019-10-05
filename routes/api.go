package routes

import (
	"youtube-manager-go/web/api" // 1

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	g := e.Group("/api") // 2
	{
		g.GET("/popular", api.FetchMostPopularVideos()) // 3
	}
}
