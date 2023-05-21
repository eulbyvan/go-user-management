package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

var dummyUsers = []entity.User{
	{
		ID:       1,
		Username: "User 1",
		Password: "Password 1",
	},
	{
		ID:       2,
		Username: "User 2",
		Password: "Password 2",
	},
}

type UserRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_InsertUser() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT INTO users").
		ExpectQuery().WithArgs(dummyUsers[0].Username, dummyUsers[0].Password).
		WillReturnRows(rows)

	repo := NewUserRepository(suite.dbMock)
	u, err := repo.InsertUser(&dummyUsers[0])
	suite.Nil(err)
	suite.Equal(int64(123), u.ID)
}

func (suite *UserRepositoryTestSuite) TestUserRepository_FindAllUser() {
	rows := sqlmock.NewRows([]string{"id", "username", "password"})
	for _, d := range dummyUsers {
		rows.AddRow(d.ID, d.Username, d.Password)
	}
	suite.sqlMock.ExpectQuery("SELECT id, username, password FROM users").
		WillReturnRows(rows)

	repo := NewUserRepository(suite.dbMock)
	users, err := repo.FindAllUser()
	suite.Nil(err)
	suite.Equal(dummyUsers, users)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
