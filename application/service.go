package service

import (
	"strings"

	"github.com/lulingua/hexa-arch/application/ports"

	domain "github.com/lulingua/hexa-arch/domain"
)

type PersonService struct {
	storage ports.Storage
}

func NewPersonaService(s ports.Storage) *PersonService {
	return &PersonService{
		storage: s,
	}
}

func (ps PersonService) ListPerson() []domain.Person {
	return ps.storage.List()
}

func (ps PersonService) AddPerson(p domain.Person) domain.Person {
	ps.storage.Add(p)
	return p
}

func UpperCaseValidator(p domain.Person) bool {

	var titleString string

	titleString = strings.Title(p.Name)

	if titleString != p.Name {
		return false
	}

	titleString = strings.Title(p.LastName)

	if titleString != p.LastName {
		return false
	}

	return true
}

func UpperCaseFixer(p domain.Person) domain.Person {

	p.Name = strings.Title(p.Name)

	p.LastName = strings.Title(p.LastName)

	return p
}
