package jsonfile

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"vita/core/entities"
)

type EntryTypesJsonRepo struct{}

func (s EntryTypesJsonRepo) loadEntryTypes() (interface{}, error) {
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

func (s EntryTypesJsonRepo) GetEntryType(entryType string) (entities.EntryType, error) {
	result := entities.EntryType{
		Key: entryType,
	}

	entryTypes, err := s.loadEntryTypes()
	if err != nil {
		return result, err
	}

	values, ok := entryTypes.(map[string]interface{})[entryType]
	if !ok {
		return result, errors.New("entry type not found")
	}

	interfaceSlice, ok := values.([]interface{})
	if !ok {
		return result, errors.New("couldn't read values for entry type")
	}

	prompts := make([]string, len(interfaceSlice))
	for i, v := range interfaceSlice {
		prompt, ok := v.(string)
		if !ok {
			return result, errors.New("failed to convert value to string")
		}
		prompts[i] = prompt
	}

	result.Prompts = prompts

	return result, nil
}

func (s EntryTypesJsonRepo) GetAllEntryTypes() ([]entities.EntryType, error) {
	return nil, errors.New("not implemented")
}

func (s EntryTypesJsonRepo) CreateEntryType(entryType entities.EntryType) error {
	return errors.New("not implemented")
}

func (s EntryTypesJsonRepo) UpdateEntryType(entryType entities.EntryType) error {
	return errors.New("not implemented")
}

func (s EntryTypesJsonRepo) DeleteEntryType(key string) error {
	return errors.New("not implemented")
}
