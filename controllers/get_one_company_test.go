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

type getCompanyTestSuite struct {
	suite.Suite
	server *server.Server

	req *http.Request
	rec *httptest.ResponseRecorder
}

func (suite *getCompanyTestSuite) SetupTest() {
	suite.server = &server.Server{Addr: "http://localhost", Port: "10000"}

	suite.req, _ = http.NewRequest(http.MethodGet, suite.server.FullUrl("company"), nil)
	suite.rec = httptest.NewRecorder()
}

func (suite *getCompanyTestSuite) AfterTest(_, _ string) {}

func (suite *getCompanyTestSuite) TestGetCompanyOk() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockCompanies(6), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	suite.req = mux.SetURLVars(suite.req, map[string]string{"companyId": "5"})

	h.GetCompany(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *getCompanyTestSuite) TestGetCompanyNotFound() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockCompanies(1), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	suite.req = mux.SetURLVars(suite.req, map[string]string{"companyId": "5"})

	h.GetCompany(suite.rec, suite.req)

	suite.Equal(http.StatusNotFound, suite.rec.Code)
}

func TestGetCompany(t *testing.T) {
	suite.Run(t, new(getCompanyTestSuite))
}
