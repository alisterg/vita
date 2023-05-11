package jsonfile

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"vita/core/entities"
)

type RoutinesJsonRepo struct{}

func (r RoutinesJsonRepo) loadRoutines() (interface{}, error) {
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

func (r RoutinesJsonRepo) GetRoutine(routine string) (entities.Routine, error) {
	result := entities.Routine{
		Key: routine,
	}

	routines, err := r.loadRoutines()
	if err != nil {
		return result, err
	}

	values, ok := routines.(map[string]interface{})[routine]
	if !ok {
		return result, errors.New("routine not found")
	}

	interfaceSlice, ok := values.([]interface{})
	if !ok {
		return result, errors.New("couldn't read values for routine")
	}

	entryTypes := make([]string, len(interfaceSlice))
	for i, v := range interfaceSlice {
		prompt, ok := v.(string)
		if !ok {
			return result, errors.New("failed to convert value to string")
		}
		entryTypes[i] = prompt
	}

	result.EntryTypes = entryTypes

	return result, nil
}

func (r RoutinesJsonRepo) GetAllRoutines() ([]entities.Routine, error) {
	return nil, errors.New("not implemented")

}

func (r RoutinesJsonRepo) CreateRoutine(routine entities.Routine) error {
	return errors.New("not implemented")
}

func (r RoutinesJsonRepo) UpdateRoutine(routine entities.Routine) error {
	return errors.New("not implemented")
}

func (r RoutinesJsonRepo) DeleteRoutine(key string) error {
	return errors.New("not implemented")
}
