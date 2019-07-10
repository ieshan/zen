package route

import (
	"github.com/ieshan/zen/conn"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	db *conn.DBClient
	h *Handler
}

func (suite *UserTestSuite) SetupTest()  {
	db := conn.GetInMemorySQLiteConn()
	suite.db = db
	suite.h = &Handler{
		db: db,
	}
}

func (suite *UserTestSuite) TearDownTest()  {
	err := suite.db.Close()
	if err != nil {
		suite.T().Log("Error closing DB connection")
	}
}

func (suite *UserTestSuite) TestUserList() {
	e := echo.New()
	Initialize(e, suite.db)
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	suite.Equal(rec.Code, http.StatusOK)
	suite.Equal("User List", rec.Body.String())
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
