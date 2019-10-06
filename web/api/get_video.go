package api

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

// VideoResponse video list
type VideoResponse struct {
	VideoList *youtube.VideoListResponse `json:"video_list"`
}

// GetVideo VideoResponse 構造体に詰めて返す
func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		videoID := c.Param("id")
		call := yts.Videos.
			List("id,snippet").
			Id(videoID)
		res, err := call.Do()

		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		v := VideoResponse{
			VideoList: res,
		}
		return c.JSON(fasthttp.StatusOK, v)
	}
}
