package use_cases

import (
	"encoding/json"
	"vita/core/entities"
	"vita/core/repositories"
)

func GetAllEntryTypesJson(entryTypeRepo repositories.EntryTypeRepository) (string, error) {
	entryTypes, err := entryTypeRepo.GetAllEntryTypes()
	if err != nil {
		return "", err
	}

	entryTypesJson, err := json.Marshal(entryTypes)
	if err != nil {
		return "", err
	}

	return string(entryTypesJson), nil
}

func CreateEntryTypeFromJson(entryTypeRepo repositories.EntryTypeRepository, entryTypeJson string) error {
	var entryType entities.EntryType
	err := json.Unmarshal([]byte(entryTypeJson), &entryType)
	if err != nil {
		return err
	}

	err2 := entryTypeRepo.CreateEntryType(entryType)
	if err2 != nil {
		return err2
	}

	return nil
}

func UpdateEntryTypeFromJson(entryTypeRepo repositories.EntryTypeRepository, entryTypeJson string) error {
	var entryType entities.EntryType
	err := json.Unmarshal([]byte(entryTypeJson), &entryType)
	if err != nil {
		return err
	}

	err2 := entryTypeRepo.UpdateEntryType(entryType)
	if err2 != nil {
		return err2
	}

	return nil
}
