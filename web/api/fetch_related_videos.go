package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

// FetchRelatedVideos 関連動画の取得
func FetchRelatedVideos() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		videoID := c.Param("id")
		call := yts.Search. // 1
					List("id,snippet").
					RelatedToVideoId(videoID). // 2
					Type("video").             // 3
					MaxResults(3)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
