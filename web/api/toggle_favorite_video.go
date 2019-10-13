package api

import (
	"youtube-manager-go/middlewares"
	"youtube-manager-go/models"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

// ToggleFavoriteVideoResponse type
type ToggleFavoriteVideoResponse struct {
	VideoID    string `json:"video_id"`
	IsFavorite bool   `json:"is_favorite"`
}

// ToggleFavoriteVideo func
func ToggleFavoriteVideo() echo.HandlerFunc {
	return func(c echo.Context) error {
		dbs := c.Get("dbs").(*middlewares.DatabaseClient)
		videoID := c.Param("id")
		token := c.Get("auth").(*auth.Token)
		user := models.User{}
		if dbs.DB.Table("users").
			Where(models.User{UID: token.UID}).First(&user).RecordNotFound() {

			user = models.User{UID: token.UID}
			dbs.DB.Create(&user)
		}

		favorite := models.Favorite{}
		isFavorite := false
		if dbs.DB.Table("favorites").
			Where(models.Favorite{UserID: user.ID, VideoID: videoID}).First(&favorite).RecordNotFound() {

			favorite = models.Favorite{UserID: user.ID, VideoID: videoID}
			dbs.DB.Create(&favorite)
			isFavorite = true
		} else {
			dbs.DB.Delete(&favorite)
		}

		res := ToggleFavoriteVideoResponse{
			VideoID:    videoID,
			IsFavorite: isFavorite,
		}

		return c.JSON(fasthttp.StatusOK, res)
	}
}
