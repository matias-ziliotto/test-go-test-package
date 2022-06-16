package mock

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
)

const (
	BASE_URL = "http://localhost:8080"
)

type Mock struct {
	Request  request  `json:"request"`
	Response response `json:"response"`
}

type request struct {
	Url               string   `json:"url"`
	Method            string   `json:"method"`
	QueryParameters   []string `json:"query_parameters"`
	Headers           struct{} `json:"headers"`
	Body              string   `json:"body"`
	ExpectedCallCount int      `json:"expected_call_count"`
}
type response struct {
	Status  int      `json:"status"`
	Headers struct{} `json:"headers"`
	Body    string   `json:"body"`
}

func (m Mock) CreateMock(filePath string) *rest.Mock {
	mock, err := readJsonFile(filePath)

	if err != nil {
		return &rest.Mock{}
	}

	restMock := mock.createRestMock()
	return restMock
}

func (m Mock) createRestMock() *rest.Mock {
	var restMock rest.Mock

	restMock.URL = m.Request.Url
	restMock.HTTPMethod = m.Request.Method
	restMock.RespHTTPCode = m.Response.Status
	restMock.RespBody = m.Response.Body
	restMock.ExpectedCallCount = m.Request.ExpectedCallCount

	// TODO: faltan mas datos por agregar, headers, query parameters, etc.

	return &restMock
}

func readJsonFile(filePath string) (Mock, error) {
	absPath, err := filepath.Abs("../../tests/mocks/" + filePath + ".json")

	if err != nil {
		return Mock{}, err
	}

	fileData, err := ioutil.ReadFile(absPath)

	if err != nil {
		return Mock{}, err
	}

	var mock Mock
	err = json.Unmarshal(fileData, &mock)

	if err != nil {
		return Mock{}, err
	}

	return mock, nil
}
