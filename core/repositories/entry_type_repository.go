package repositories

import "vita/core/entities"

type EntryTypeRepository interface {
	GetEntryType(key string) (entities.EntryType, error)
	GetAllEntryTypes() ([]entities.EntryType, error)
	CreateEntryType(entryType entities.EntryType) error
	UpdateEntryType(entities.EntryType) error
	DeleteEntryType(key string) error
}
