package entities

type Routine struct {
	Key        string   `json:"key"`        // dynamo column: key
	EntryTypes []string `json:"entryTypes"` // dynamo column: entry_types
}
