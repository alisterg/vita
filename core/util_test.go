package core

import (
	"testing"

	"vita/core/entities"
)

func compareKvPairs(a, b []KvPair) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestMapSorter(t *testing.T) {
	// Test case 1: Empty input map
	inputMap1 := make(map[string]string)
	expectedResult1 := []KvPair{}
	result1 := MapSorter(inputMap1)
	if len(result1) != len(expectedResult1) {
		t.Errorf("Expected %d items, but got %d", len(expectedResult1), len(result1))
	}

	// Test case 2: Map with one key-value pair
	inputMap2 := map[string]string{
		"key1": "value1",
	}
	expectedResult2 := []KvPair{
		{"key1", "value1"},
	}
	result2 := MapSorter(inputMap2)
	if len(result2) != len(expectedResult2) {
		t.Errorf("Expected %d items, but got %d", len(expectedResult2), len(result2))
	} else {
		if !compareKvPairs(result2, expectedResult2) {
			t.Errorf("Expected %v, but got %v", expectedResult2, result2)
		}
	}

	// Test case 3: Map with multiple key-value pairs
	inputMap3 := map[string]string{
		"key2": "value2",
		"key1": "value1",
		"key3": "value3",
	}
	expectedResult3 := []KvPair{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
	}
	result3 := MapSorter(inputMap3)
	if len(result3) != len(expectedResult3) {
		t.Errorf("Expected %d items, but got %d", len(expectedResult3), len(result3))
	} else {
		if !compareKvPairs(result3, expectedResult3) {
			t.Errorf("Expected %v, but got %v", expectedResult3, result3)
		}
	}
}

func TestSearchEntries(t *testing.T) {
	// Test case 1: Empty entries slice
	entries1 := []entities.Entry{}
	query1 := "test"
	expectedResult1 := []entities.Entry{}
	result1 := SearchEntries(entries1, query1)
	if len(result1) != len(expectedResult1) {
		t.Errorf("Expected %d items, but got %d", len(expectedResult1), len(result1))
	}

	// Test case 2: No matching entries
	entries2 := []entities.Entry{
		{Data: make(map[string]string)},
		{Data: make(map[string]string)},
	}
	query2 := "test"
	expectedResult2 := []entities.Entry{}
	result2 := SearchEntries(entries2, query2)
	if len(result2) != len(expectedResult2) {
		t.Errorf("Expected %d items, but got %d", len(expectedResult2), len(result2))
	}

	// Test case 3: Matching entries
	entries3 := []entities.Entry{
		{Data: map[string]string{"key1": "test1", "key2": "data2"}},
		{Data: map[string]string{"key3": "data3", "key4": "test2"}},
		{Data: map[string]string{"key5": "data4", "key6": "data5"}},
	}
	query3 := "test"
	expectedResult3 := []entities.Entry{
		{Data: map[string]string{"key1": "test1", "key2": "data2"}},
		{Data: map[string]string{"key3": "data3", "key4": "test2"}},
	}
	result3 := SearchEntries(entries3, query3)
	if len(result3) != len(expectedResult3) {
		t.Errorf("Expected %d items, but got %d", len(expectedResult3), len(result3))
	} else {
		for i := range expectedResult3 {
			for k, v := range expectedResult3[i].Data {
				if result3[i].Data[k] != v {
					t.Errorf("Expected %v, but got %v", expectedResult3, result3)
					break
				}
			}
		}
	}
}
