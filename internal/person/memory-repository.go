package person

import (
	"strings"
)

type MemoryRepository struct {
	persons []Person
}

var memoryRepository *MemoryRepository

func NewMemoryRepository() *MemoryRepository {
	if memoryRepository == nil {
		memoryRepository = &MemoryRepository{
			persons: []Person{},
		}
	}

	return memoryRepository
}

func (m *MemoryRepository) FindById(id int) (person Person, err error) {
	for _, p := range m.persons {
		if p.Id == id {
			person = p
			break
		}
	}

	return person, nil
}

func (m *MemoryRepository) FindByTerm(t string) (persons []Person, err error) {
	for _, p := range m.persons {
		s := p.Nome + " " + p.Apelido + " " + strings.Join(p.Stack, ",")
		if strings.Contains(s, t) {
			persons = append(persons, p)
		}
	}

	return persons, nil
}

func (m *MemoryRepository) Count() (int, error) {
	return len(m.persons), nil
}

func (m *MemoryRepository) Insert(p Person) (id int, err error) {
	p.Id = len(m.persons) + 1
	m.persons = append(m.persons, p)

	return p.Id, nil
}
