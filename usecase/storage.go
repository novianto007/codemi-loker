package usecase

import (
	"codemi/loker/models"
	"codemi/loker/structure"
	"errors"
)

type Storage struct {
	Locker    []*models.Data
	Available *structure.OrderedList
	Index     *structure.Index
}

func NewStorage(size int) *Storage {
	available := &structure.OrderedList{}
	locker := []*models.Data{}
	for i := 0; i < size; i++ {
		locker = append(locker, nil)
		available.Push(i)
	}
	return &Storage{
		Locker:    locker,
		Available: available,
		Index:     structure.NewIndex(),
	}
}

func (s Storage) Save(identityType string, number int64) (int, error) {
	data := models.Data{Type: identityType, Number: number}
	index, err := s.Available.Pop()
	if err != nil {
		return index, errors.New("Maaf loker sudah penuh")
	}
	err = s.Index.Save(number, index)
	if err != nil {
		s.Available.Push(index)
		return 0, err
	}
	s.Locker[index] = &data
	return index, nil
}

func (s Storage) GetAll() []*models.Data {
	return s.Locker
}

func (s Storage) Get(number int64) (int, error) {
	index, err := s.Index.Get(number)
	if err != nil {
		return 0, err
	}
	return index, nil
}

func (s Storage) GetByType(identityType string) []*models.Data {
	data := []*models.Data{}
	for _, row := range s.Locker {
		if row.Type == identityType {
			data = append(data, row)
		}
	}
	return data
}

func (s *Storage) Remove(index int) error {
	if index < 0 || index >= len(s.Locker) {
		return errors.New("Nomer loker tidak valid")
	}
	if s.Locker[index] == nil {
		return errors.New("Loker sudah kosong")
	}
	data := s.Locker[index]
	s.Index.Remove(data.Number)
	s.Locker[index] = nil
	s.Available.Push(index)
	return nil
}
