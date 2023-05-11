package entities

type EntryType struct {
	Key     string   // dynamo column: key
	Prompts []string // dynamo column: prompts
}
