package route

import (
	"github.com/ieshan/zen/conn"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) UserList(c echo.Context) error {
	return c.String(http.StatusOK, "User List")
}

func (h *Handler) UserDetail(c echo.Context) error {
	return c.String(http.StatusOK, "User Detail")
}

func InitUserRoutes(e *echo.Group, db *conn.DBClient) {
	groupPrefix := "user:"
	ur := &Handler{
		db,
	}
	e.GET("", ur.UserList).Name = groupPrefix + "list"
	//e.GET("", ur.UserList).Name = groupPrefix + "detail"
}
