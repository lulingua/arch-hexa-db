package mysql

import (
	"database/sql"

	"github.com/lulingua/hexa-arch/domain"
)

type PersonMySQL struct {
	conn *sql.DB
}

func NewPersonMySQL(conn *sql.DB) *PersonMySQL {
	return &PersonMySQL{conn}
}

func (m *PersonMySQL) Add(p domain.Person) {
}

func (m *PersonMySQL) List() []domain.Person {
	return nil
}

/*
func (m *PersonMySQL) Exists(uuid string) bool {
	return true
}*/
