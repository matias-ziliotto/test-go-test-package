package mock

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Mock struct {
	Request  request  `json:"request"`
	Response response `json:"response"`
}

type request struct {
	URL               string      `json:"url"`
	Method            string      `json:"method"`
	QueryParameters   struct{}    `json:"query_parameters"`
	Headers           struct{}    `json:"headers"`
	IgnoreExtraFields bool        `json:"ignore_extra_fields"`
	Body              interface{} `json:"body"`
	ExpectedCallCount int64       `json:"expected_call_count"`
}
type response struct {
	Status  int      `json:"status"`
	Headers struct{} `json:"headers"`
	Body    struct{} `json:"body"`
}

func ReadJsonFile(filePath string) (Mock, error) {
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
