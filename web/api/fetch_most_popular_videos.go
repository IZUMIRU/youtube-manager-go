package api // 1

import (
	"context"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// FetchMostPopularVideos Youtubeの人気動画を取得する
func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		key := os.Getenv("API_KEY") // 1

		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key)) // 2
		if err != nil {
			logrus.Fatalf("Error creating new Youtube service: %v", err)
		}

		call := yts.Videos. // 3
					List("id,snippet").
					Chart("mostPopular").
					MaxResults(3)

		pageToken := c.QueryParam("pageToken")
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}

		res, err := call.Do() // 4
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
