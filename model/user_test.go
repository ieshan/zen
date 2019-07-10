package model

import (
	"github.com/ieshan/zen/conn"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	db *conn.DBClient
	testUser1 *User
	testUser2 *User
}

func (suite *UserTestSuite) SetupTest()  {
	suite.db = conn.GetInMemorySQLiteConn()
	Migrate(suite.db)
	user1 := User{
		Name: "Test User 1",
		Email: "email1@test.tld",
	}
	user1.Create(suite.db)
	user2 := User{
		Name: "Test User 2",
		Email: "email2@test.tld",
	}
	user2.Create(suite.db)
	suite.testUser1 = &user1
	suite.testUser2 = &user2
}

func (suite *UserTestSuite) TearDownTest()  {
	err := suite.db.Close()
	if err != nil {
		suite.T().Log("Error closing DB connection")
	}
}

func (suite *UserTestSuite) TestUserCreate() {
	user := User{
		Name: "Test",
		Email: "email@test.tld",
	}
	suite.Equal(true, suite.db.NewRecord(user))
	user.Create(suite.db)
	suite.Equal(false, suite.db.NewRecord(user))
}

func (suite *UserTestSuite) TestUserFetch() {
	user := User{}
	user.Get(suite.db, "id=?", suite.testUser1.ID)
	suite.Equal(suite.testUser1.ID, user.ID)
	suite.Equal(suite.testUser1.Name, user.Name)
}

func (suite *UserTestSuite) TestUserUpdate() {
	editedName := "Edited Name"
	suite.testUser1.Name = editedName
	suite.testUser1.Update(suite.db)
	user := User{}
	user.Get(suite.db, "id=?", suite.testUser1.ID)
	suite.Equal(editedName, user.Name)
}

func (suite *UserTestSuite) TestUserList() {
	userList := UserList{}
	userList.Find(suite.db, 0, 10, "")
	suite.Equal(2, len(userList))
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
