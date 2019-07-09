package model

import (
	"github.com/ieshan/zen/conn"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	db *conn.DBClient
}

func (suite *UserTestSuite) SetupTest()  {
	suite.db = conn.GetInMemorySQLiteConn()
	Migrate(suite.db)
}

func (suite *UserTestSuite) TearDownTest()  {
	err := suite.db.Close()
	if err != nil {
		suite.T().Log("Error closing DB connection")
	}
}

func (suite *UserTestSuite) TestUser() {
	user1 := User{
		Name: "Test",
		Email: "email@test.tld",
	}
	user1.Create(suite.db)
	suite.Equal(suite.db.NewRecord(user1), false)
	user1.Name = "Edited name"
	user1.Update(suite.db)
	user2 := User{}
	user2.Get(suite.db, "id=?", user1.ID)
	suite.Equal(user2.Name, "Edited name")
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
