package structure

import "testing"

var indexed map[int64]int = map[int64]int{1001: 1, 1002: 2}

func TestSaveIndex(t *testing.T) {
	var key int64 = 1003
	val := 3
	index := setupDefaultData()
	index.Save(key, val)
	v, err := index.Get(key)
	if err != nil {
		t.Fatal("Error shoud not be appear")
	}
	if v != val {
		t.Fatal("Get wrong value")
	}
}

func TestGetIndex(t *testing.T) {
	var validKey int64 = 1002
	var invalidKey int64 = 1003
	index := setupDefaultData()
	v, err := index.Get(validKey)
	if err != nil {
		t.Fatal("Error shoud not be appear")
	}
	if v != indexed[validKey] {
		t.Fatal("Get wrong value")
	}
	_, err = index.Get(invalidKey)
	if err == nil {
		t.Fatal("invalid key must give error")
	}
}

func TestRemoveIndex(t *testing.T) {
	var key int64 = 1002
	index := setupDefaultData()
	err := index.Remove(key)
	if err != nil {
		t.Fatal("Error shoud not be appear")
	}
	err = index.Remove(key)
	if err == nil {
		t.Fatal("remove same key must give error")
	}
}

func setupDefaultData() *Index {
	index := NewIndex()
	for key, val := range indexed {
		index.Save(key, val)
	}
	return index
}
