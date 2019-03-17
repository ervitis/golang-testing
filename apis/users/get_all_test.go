package users

import (
	"encoding/json"
	"github.com/ervitis/golang-testing/server"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetUsersTestSuite struct {
	suite.Suite
	server *server.Server

	req *http.Request
	rec *httptest.ResponseRecorder
}

func (suite *GetUsersTestSuite) SetupTest() {
	suite.server = &server.Server{Addr: "http://localhost", Port: "10000"}

	suite.req, _ = http.NewRequest(http.MethodGet, suite.server.Url(), nil)
	suite.rec = httptest.NewRecorder()
}

func (suite *GetUsersTestSuite) AfterTest(_, _ string) {}

func (suite *GetUsersTestSuite) TestGetAllUsersOk() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockUsers(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *GetUsersTestSuite) TestGetAllUsersNoPage() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockUsers(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *GetUsersTestSuite) TestGetAllUsersWithPage() {
	q := suite.req.URL.Query()
	q.Add("page", "2")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockUsers(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)

	var resp []*User
	if err := json.NewDecoder(suite.rec.Body).Decode(&resp); err != nil {
		panic(err)
	}

	suite.Equal(15, len(resp))
	suite.Equal(16, resp[0].Id)
}

func (suite *GetUsersTestSuite) TestGetAllUsersWithWrongPageNumber() {
	q := suite.req.URL.Query()
	q.Add("page", "ab")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockUsers(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusBadRequest, suite.rec.Code)
}

func (suite *GetUsersTestSuite) TestGetAllUsersPageIsBigThanElements() {
	q := suite.req.URL.Query()
	q.Add("page", "40")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockUsers(3), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)

	var resp []*User
	if err := json.NewDecoder(suite.rec.Body).Decode(&resp); err != nil {
		panic(err)
	}

	suite.Equal(0, len(resp))
	suite.Equal([]*User{}, resp)
}

func (suite *GetUsersTestSuite) TestGetAllUsersPageIEqualThanElements() {
	q := suite.req.URL.Query()
	q.Add("page", "2")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockUsers(30), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func TestGetUsers(t *testing.T) {
	suite.Run(t, new(GetUsersTestSuite))
}
