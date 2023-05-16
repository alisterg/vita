package repositories

import "vita/core/entities"

type EntryRepository interface {
	CreateEntry(entry entities.Entry) error
	UpdateEntry(entry entities.Entry) error
	BulkCreateEntries(entries []entities.Entry) error
	GetAllEntries() ([]entities.Entry, error)
	GetAllEntriesForType(entryType string) ([]entities.Entry, error)
}
