package instance

import (
	"github.com/ieshan/zen/logger"
	"github.com/labstack/echo"
	"sync"
)

var echoOnce sync.Once
var echoInstance *echo.Echo

func GetEcho() *echo.Echo {
	echoOnce.Do(func() {
		echoInstance = echo.New()
		echoInstance.Logger = logger.GetLogger()
	})
	return echoInstance
}
