package controllers

import (
	"github.com/ervitis/golang-testing/server"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type getUserTestSuite struct {
	suite.Suite
	server *server.Server

	req *http.Request
	rec *httptest.ResponseRecorder
}

func (suite *getUserTestSuite) SetupTest() {
	suite.server = &server.Server{Addr: "http://localhost", Port: "10000"}

	suite.req, _ = http.NewRequest(http.MethodGet, suite.server.FullUrl("user"), nil)
	suite.rec = httptest.NewRecorder()
}

func (suite *getUserTestSuite) AfterTest(_, _ string) {}

func (suite *getUserTestSuite) TestGetUserOk() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockUsers(6), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	suite.req = mux.SetURLVars(suite.req, map[string]string{"userId": "5"})

	h.GetUser(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *getUserTestSuite) TestGetUserNotFound() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockUsers(1), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	suite.req = mux.SetURLVars(suite.req, map[string]string{"userId": "5"})

	h.GetUser(suite.rec, suite.req)

	suite.Equal(http.StatusNotFound, suite.rec.Code)
}

func TestGetUser(t *testing.T) {
	suite.Run(t, new(getUserTestSuite))
}
