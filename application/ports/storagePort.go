package ports

import "github.com/lulingua/hexa-arch/domain"

type Storage interface {
	//Exists(uuid string) bool
	Add(p domain.Person)
	List() []domain.Person
}
