package api

import (
	"youtube-manager-go/middlewares"
	"youtube-manager-go/models"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/youtube/v3"
)

// VideoResponse type
type VideoResponse struct {
	VideoList  *youtube.VideoListResponse `json:"video_list"`
	IsFavorite bool                       `json:"is_favorite"`
}

// GetVideo func
func GetVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		yts := c.Get("yts").(*youtube.Service)
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		token := c.Get("auth").(*auth.Token)

		videoID := c.Param("id")

		isFavorite := false
		if token != nil {
			favorite := models.Favorite{}
			isFavoriteNotFound := dbs.DB.Table("favorites").
				Joins("INNER JOIN users ON users.id = favorites.user_id").
				Where(models.User{UID: token.UID}).
				Where(models.Favorite{VideoID: videoID}).
				First(&favorite).
				RecordNotFound()

			logrus.Debug("isFavoriteNotFound: ", isFavoriteNotFound)
			if !isFavoriteNotFound {
				isFavorite = true
			}
		}

		call := yts.Videos.
			List("id,snippet").
			Id(videoID)

		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}

		v := VideoResponse{
			VideoList:  res,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, v)
	}
}
