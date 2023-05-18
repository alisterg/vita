package entities

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	Uuid      string            // dynamo column: uuid
	EntryType string            // dynamo column: type
	Data      map[string]string // dynamo column: data
	CreatedAt int64             // dynamo column: date
}

type EntryDto struct {
	EntryType string            `json:"entryType"`
	EntryData map[string]string `json:"entryData"`
}

func EntryFactory(entryType string, entryData map[string]string) Entry {
	return Entry{
		Uuid:      uuid.New().String(),
		EntryType: entryType,
		Data:      entryData,
		CreatedAt: time.Now().Unix(),
	}
}
