package route

import (
	"github.com/labstack/echo"
	"net/http"
)

func DummyRoutes(e *echo.Group) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
}
