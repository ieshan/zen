package instance

import (
	"github.com/ieshan/zen/config"
	"github.com/ieshan/zen/conn"
	"sync"
)

var mysqlOnce sync.Once
var mysqlConn *conn.DBClient

func GetDB() *conn.DBClient {
	mysqlOnce.Do(func() {
		mysqlConn = conn.GetMySQLConn(config.GetDB())
	})
	return mysqlConn
}
