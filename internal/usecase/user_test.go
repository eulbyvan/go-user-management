package usecase

import (
	"errors"
	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
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

type UserRepositoryMock struct {
	mock.Mock
}

func (u *UserRepositoryMock) InsertUser(user *entity.User) (*entity.User, error) {
	args := u.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), nil
}

func (u *UserRepositoryMock) UpdateUser(*entity.User) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepositoryMock) DeleteUser(*entity.User) error {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepositoryMock) FindUserByID(int64) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepositoryMock) FindUserByUsername(string) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *UserRepositoryMock) FindAllUser() ([]entity.User, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.User), nil
}

type UserUsecaseTestSuite struct {
	suite.Suite
	repoMock *UserRepositoryMock
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(UserRepositoryMock)
}

func (suite *UserUsecaseTestSuite) TestInsertUser() {
	user := dummyUsers[0]
	suite.repoMock.On("InsertUser", &user).Return(&user, nil)
	userUsecase := NewUserUsecase(suite.repoMock)
	result, err := userUsecase.InsertUser(&user)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &user, result)
}

func (suite *UserUsecaseTestSuite) TestInsertUserError() {
	user := dummyUsers[0]
	suite.repoMock.On("InsertUser", &user).Return(nil, errors.New("error"))
	userUsecase := NewUserUsecase(suite.repoMock)
	result, err := userUsecase.InsertUser(&user)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *UserUsecaseTestSuite) TestFindAllUser() {
	suite.repoMock.On("FindAllUser").Return(dummyUsers, nil)
	userUsecase := NewUserUsecase(suite.repoMock)
	result, err := userUsecase.FindAllUser()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyUsers, result)
}

func (suite *UserUsecaseTestSuite) TestFindAllUserError() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current working directory:", dir)

	suite.repoMock.On("FindAllUser").Return(nil, errors.New("error"))
	userUsecase := NewUserUsecase(suite.repoMock)
	result, err := userUsecase.FindAllUser()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
