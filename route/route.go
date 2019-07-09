package route

import (
	"github.com/ieshan/zen/conn"
	"github.com/labstack/echo"
)

type Handler struct {
	db *conn.DBClient
}

func Initialize(e *echo.Echo, db *conn.DBClient)  {
	InitUserRoutes(e.Group("/users"), db)
}
