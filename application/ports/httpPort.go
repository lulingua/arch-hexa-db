package ports

import (
	"github.com/lulingua/hexa-arch/domain"
)

type PersonService interface {
	//GetPerson() domain.Person
	ListPerson() []domain.Person
	AddPerson(p domain.Person) domain.Person
}
