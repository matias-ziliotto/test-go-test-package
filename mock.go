package mock

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ReadJsonFile(filePath string, data interface{}) error {
	absPath, err := filepath.Abs("../../tests/mocks/" + filePath + ".json")
	log.Println("Path:", err)
	log.Println("Path:", absPath)
	fileData, err := ioutil.ReadFile(absPath)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}
