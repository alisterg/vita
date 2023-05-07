package core

import (
	"sort"
	"strings"

	"vita/core/entities"
)

type KvPair struct {
	Key   string
	Value string
}

func MapSorter(inputMap map[string]string) []KvPair {
	var kvPairs []KvPair
	for k, v := range inputMap {
		kvPairs = append(kvPairs, KvPair{k, v})
	}

	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i].Key < kvPairs[j].Key
	})

	return kvPairs
}

func SearchEntries(entries []entities.Entry, query string) []entities.Entry {
	var result []entities.Entry

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
