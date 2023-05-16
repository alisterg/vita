package use_cases

import (
	"encoding/json"
	"vita/core/repositories"
)

func GetAllEntriesJson(entryRepo repositories.EntryRepository) (string, error) {
	entries, err := entryRepo.GetAllEntries()
	if err != nil {
		return "", err
	}

	entriesJson, err := json.Marshal(entries)
	if err != nil {
		return "", err
	}

	return string(entriesJson), nil
}
