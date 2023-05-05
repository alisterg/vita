package core

import "sort"

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
