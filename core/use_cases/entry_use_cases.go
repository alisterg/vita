package use_cases

import (
	"encoding/json"
	"vita/core/entities"
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

func CreateEntryFromJson(entryRepo repositories.EntryRepository, entryJson string) error {
	var entryDto entities.EntryDto
	err := json.Unmarshal([]byte(entryJson), &entryDto)
	if err != nil {
		return err
	}

	entry := entities.EntryFactory(entryDto.EntryType, entryDto.EntryData)

	err2 := entryRepo.CreateEntry(entry)
	if err2 != nil {
		return err2
	}

	return nil
}

func CreateBulkEntriesFromJson(entryRepo repositories.EntryRepository, entriesJson string) error {
	var entryDtos []entities.EntryDto
	err := json.Unmarshal([]byte(entriesJson), &entryDtos)
	if err != nil {
		return err
	}

	var entries []entities.Entry
	for _, entryDto := range entryDtos {
		entry := entities.EntryFactory(entryDto.EntryType, entryDto.EntryData)
		entries = append(entries, entry)
	}

	err2 := entryRepo.BulkCreateEntries(entries)
	if err2 != nil {
		return err2
	}

	return nil
}

func UpdateEntryFromJson(entryRepo repositories.EntryRepository, entryJson string) error {
	var entryDto entities.EntryDto
	err := json.Unmarshal([]byte(entryJson), &entryDto)
	if err != nil {
		return err
	}

	// TODO: this is assigning a new uuid, how is it supposed to update an existing one

	entry := entities.EntryFactory(entryDto.EntryType, entryDto.EntryData)

	err2 := entryRepo.UpdateEntry(entry)
	if err2 != nil {
		return err2
	}

	return nil

}
