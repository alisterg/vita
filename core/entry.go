package core

import "strings"

type Entry struct {
	Uuid      string            // uuid
	EntryType string            // type
	Data      map[string]string // data
	CreatedAt int64             // date
}

func SearchEntries(entries []*Entry, query string) []*Entry {
	var result []*Entry

	for _, entry := range entries {
		for _, data := range entry.Data {
			if strings.Contains(strings.ToLower(data), strings.ToLower(query)) {
				result = append(result, entry)
				break
			}
		}
	}

	return result
}
