package users

import (
	"github.com/ervitis/golang-testing/routes"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetUsersTestSuite struct {
	suite.Suite
	server *routes.Server

	req *http.Request
	rec *httptest.ResponseRecorder
}

func (suite *GetUsersTestSuite) SetupTest() {
	suite.server = &routes.Server{Addr: "http://localhost", Port: "10000"}

	suite.req, _ = http.NewRequest(http.MethodGet, suite.server.Url(), nil)
	suite.rec = httptest.NewRecorder()
}

func (suite *GetUsersTestSuite) AfterTest(_, _ string) {

}

func (suite *GetUsersTestSuite) TestGetAllUsersOk() {
	mockito := new(mockGetUsers)

	var users []*User

	mockito.On("ReadData", mock.Anything, &users).Return(nil)

	h := GetHandler{
		Reader: mockito,
	}

	h.GetAllUsers(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func TestGetUsers(t *testing.T) {
	suite.Run(t, new(GetUsersTestSuite))
}
