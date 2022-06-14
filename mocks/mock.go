package mock

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJsonFile(filePath string, data interface{}) error {
	fileData, err := ioutil.ReadFile("/tests/mocks/" + filePath + ".json")

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}
