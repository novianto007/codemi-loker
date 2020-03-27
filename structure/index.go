package structure

import "errors"

type Index struct {
	indexed map[int64]int
}

func NewIndex() *Index {
	return &Index{
		indexed: map[int64]int{},
	}
}

func (i *Index) Save(key int64, value int) error {
	if _, ok := i.indexed[key]; ok {
		return errors.New("Duplikasi nomor identitas")
	}
	i.indexed[key] = value
	return nil
}

func (i Index) Get(key int64) (int, error) {
	if val, ok := i.indexed[key]; ok {
		return val, nil
	}
	return 0, errors.New("Nomor identitas tidak ditemukan")
}

func (i *Index) Remove(key int64) error {
	if _, ok := i.indexed[key]; ok {
		delete(i.indexed, key)
		return nil
	}
	return errors.New("Nomer identitas tidak ditemukan")
}
