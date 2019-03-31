package controllers

import (
	"encoding/json"
	"github.com/ervitis/golang-testing/server"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestGetCompaniesSuite struct {
	suite.Suite
	server *server.Server

	req *http.Request
	rec *httptest.ResponseRecorder
}

func (suite *TestGetCompaniesSuite) SetupTest() {
	suite.server = &server.Server{Addr: "http://localhost", Port: "10000"}

	suite.req, _ = http.NewRequest(http.MethodGet, suite.server.Url(), nil)
	suite.rec = httptest.NewRecorder()
}

func (suite *TestGetCompaniesSuite) AfterTest(_, _ string) {}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesOk() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockCompanies(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesNoPage() {
	mockito := new(mocker)

	mockito.On("ReadData", mock.Anything).Return(mockCompanies(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesWithPage() {
	q := suite.req.URL.Query()
	q.Add("page", "2")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)

	var resp []*company
	if err := json.NewDecoder(suite.rec.Body).Decode(&resp); err != nil {
		panic(err)
	}

	suite.Equal(15, len(resp))
	suite.Equal(16, resp[0].Id)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesWithWrongPageNumber() {
	q := suite.req.URL.Query()
	q.Add("page", "ab")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusBadRequest, suite.rec.Code)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesPageIsBigThanElements() {
	q := suite.req.URL.Query()
	q.Add("page", "40")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(3), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)

	var resp []*company
	if err := json.NewDecoder(suite.rec.Body).Decode(&resp); err != nil {
		panic(err)
	}

	suite.Equal(0, len(resp))
	suite.Equal([]*company{}, resp)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesPageIEqualThanElements() {
	q := suite.req.URL.Query()
	q.Add("page", "2")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(30), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesFilteredByUser() {
	q := suite.req.URL.Query()
	q.Add("userId", "2")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(30), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesFilteredByUserNotFound() {
	q := suite.req.URL.Query()
	q.Add("userId", "4")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(3), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusNotFound, suite.rec.Code)
}

func (suite *TestGetCompaniesSuite) TestGetAllCompaniesFilteredByUserOneElement() {
	q := suite.req.URL.Query()
	q.Add("userId", "4")
	suite.req.URL.RawQuery = q.Encode()

	mockito := new(mocker)
	mockito.On("ReadData", mock.Anything).Return(mockCompanies(5), nil)

	h := ReqHandler{
		Reader: mockito,
	}

	h.GetAllCompanies(suite.rec, suite.req)

	suite.Equal(http.StatusOK, suite.rec.Code)

	var resp []*company
	if err := json.NewDecoder(suite.rec.Body).Decode(&resp); err != nil {
		panic(err)
	}

	suite.Equal(1, len(resp))
}

func TestGetCompanies(t *testing.T) {
	suite.Run(t, new(TestGetCompaniesSuite))
}