package sliceinmemory

import "github.com/lulingua/hexa-arch/domain"

type InMemoryDB struct {
	Persons []domain.Person
}

func NewInMemoryDB(p []domain.Person) *InMemoryDB {
	return &InMemoryDB{
		Persons: p,
	}
}

func (i *InMemoryDB) Add(p domain.Person) {
	i.Persons = append(i.Persons, p)
}

func (i *InMemoryDB) List() []domain.Person {
	return i.Persons
}
