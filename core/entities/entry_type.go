package entities

type EntryType struct {
	Key     string   `json:"key"`     // dynamo column: key
	Prompts []string `json:"prompts"` // dynamo column: prompts
}
