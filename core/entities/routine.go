package entities

type Routine struct {
	Key        string   // dynamo column: key
	EntryTypes []string // dynamo column: entry_types
}
