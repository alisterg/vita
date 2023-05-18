package use_cases

import (
	"encoding/json"
	"vita/core/entities"
	"vita/core/repositories"
)

func GetAllRoutinesJson(routinesRepo repositories.RoutineRepository) (string, error) {
	routines, err := routinesRepo.GetAllRoutines()
	if err != nil {
		return "", err
	}

	routinesJson, err := json.Marshal(routines)
	if err != nil {
		return "", err
	}

	return string(routinesJson), nil
}

func CreateRoutineFromJson(routineRepo repositories.RoutineRepository, routineJson string) error {
	var routine entities.Routine
	err := json.Unmarshal([]byte(routineJson), &routine)
	if err != nil {
		return err
	}

	err2 := routineRepo.CreateRoutine(routine)
	if err2 != nil {
		return err2
	}

	return nil
}

func UpdateRoutineFromJson(routineRepo repositories.RoutineRepository, routineJson string) error {
	var routine entities.Routine
	err := json.Unmarshal([]byte(routineJson), &routine)
	if err != nil {
		return err
	}

	err2 := routineRepo.UpdateRoutine(routine)
	if err2 != nil {
		return err2
	}

	return nil
}

func DeleteRoutine(routineRepo repositories.RoutineRepository, key string) error {
	err := routineRepo.DeleteRoutine(key)
	if err != nil {
		return err
	}

	return nil
}
