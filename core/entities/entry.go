package entities

type Entry struct {
	Uuid      string            // dynamo column: uuid
	EntryType string            // dynamo column: type
	Data      map[string]string // dynamo column: data
	CreatedAt int64             // dynamo column: date
}
