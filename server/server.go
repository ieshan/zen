package server

import (
	"github.com/ieshan/zen/logger"
	"github.com/ieshan/zen/route"
	"github.com/labstack/echo"
	"strconv"
)

func ServeHTTP(port int) {
	e := echo.New()
	route.DummyRoutes(e.Group("/user"))
	log := logger.GetLogger()
	log.ErrorFatal("ServeHTTP error.", e.Start(":" + strconv.Itoa(int(port))))
}
