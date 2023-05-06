package app

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func loadRoutines() (interface{}, error) {
	file, err := os.Open("routines.json")
	if err != nil {
		return nil, errors.New("couldn't load routines.json")
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("couldn't read routines.json")
	}

	var result interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, errors.New("couldn't understand routines.json")
	}

	return result, nil
}

func LoadPromptsForRoutine(routine string) ([]interface{}, error) {
	routines, err := loadRoutines()
	if err != nil {
		return nil, err
	}

	values, ok := routines.(map[string]interface{})[routine]
	if !ok {
		return nil, errors.New("routine not found")
	}

	result, ok := values.([]interface{})
	if !ok {
		return nil, errors.New("couldn't read values for routine")
	}

	return result, nil
}
