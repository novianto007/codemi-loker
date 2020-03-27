package usecase

import "testing"

func TestSave(t *testing.T) {
	storage := NewStorage(1)
	index, err := storage.Save("ktp", int64(1001))
	if err != nil {
		t.Fail("error should nil")
	}
	if index != 0 {
		t.Fatal("wrong index number")
	}
	_, err = storage.Save("ktp", int64(1001))
	if err == nil {
		t.Fail("should give an error")
	}
}
