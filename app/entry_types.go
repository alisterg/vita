package app

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func loadEntryTypes() (interface{}, error) {
	file, err := os.Open("entry_types.json")
	if err != nil {
		return nil, errors.New("couldn't load entry_types.json")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("couldn't read entry_types.json")
	}

	var result interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, errors.New("couldn't understand entry_types.json")
	}

	return result, nil
}

func LoadPrompts(entryType string) ([]interface{}, error) {
	entryTypes, err := loadEntryTypes()
	if err != nil {
		return nil, err
	}

	values, ok := entryTypes.(map[string]interface{})[entryType]
	if !ok {
		return nil, errors.New("entry type not found")
	}

	result, ok := values.([]interface{})
	if !ok {
		return nil, errors.New("couldn't read values for entry type")
	}

	return result, nil
}