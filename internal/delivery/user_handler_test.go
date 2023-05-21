package delivery

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eulbyvan/go-user-management/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockResponse struct {
	Data entity.User `json:"data"`
}

var dummyUsers = []entity.User{
	{
		ID:       0,
		Username: "User 0",
		Password: "Password 0",
	},
	{
		ID:       1,
		Username: "User 1",
		Password: "Password 1",
	},
}

type userUseCaseMock struct {
	mock.Mock
}

func (u *userUseCaseMock) InsertUser(user *entity.User) (*entity.User, error) {
	args := u.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), nil
}

func (u *userUseCaseMock) UpdateUser(user *entity.User) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *userUseCaseMock) DeleteUser(user *entity.User) error {
	// TODO implement me
	panic("implement me")
}

func (u *userUseCaseMock) FindUserByID(i int64) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *userUseCaseMock) FindUserByUsername(s string) (*entity.User, error) {
	args := u.Called(s)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), nil
}

func (u *userUseCaseMock) FindAllUser() ([]entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u *userUseCaseMock) Login(user *entity.User) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

type UserHandlerTestSuite struct {
	suite.Suite
	useCaseMock     *userUseCaseMock
	routerTest      *gin.Engine
	routerGroupTest *gin.RouterGroup
}

func (suite *UserHandlerTestSuite) SetupTest() {
	suite.useCaseMock = new(userUseCaseMock)
	suite.routerTest = gin.Default()
}

func (suite *UserHandlerTestSuite) Test_InsertUser_Success() {
	dummyUser := dummyUsers[0]
	suite.useCaseMock.On("FindUserByUsername", dummyUser.Username).Return(nil, nil)
	suite.useCaseMock.On("InsertUser", &dummyUser).Return(&dummyUser, nil)
	userHandler := NewUserHandler(suite.useCaseMock)
	insertUserHandler := userHandler.InsertUser
	suite.routerTest.POST("/", insertUserHandler)

	rr := httptest.NewRecorder()
	reqBody, _ := json.Marshal(dummyUser)
	request, _ := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	request.Header.Set("Content-Type", "application/json")
	suite.routerTest.ServeHTTP(rr, request)
	assert.Equal(suite.T(), rr.Code, 201)

	a := rr.Body.String()
	actualUser := new(MockResponse)
	json.Unmarshal([]byte(a), actualUser)
	assert.Equal(suite.T(), dummyUsers[0].Username, actualUser.Data.Username)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
