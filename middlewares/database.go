package middlewares

import (
	"youtube-manager-go/databases"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// DatabaseClient struct
type DatabaseClient struct {
	DB *gorm.DB
}

// DatabaseService MiddlewareFunc
func DatabaseService() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := databases.Connect()
			d := DatabaseClient{DB: session}

			defer d.DB.Close()

			// output sql query
			d.DB.LogMode(true)

			c.Set("dbs", &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
