package person

import (
	"fmt"
	"strings"
	"time"

	"github.com/backeseduardo/rinha-backend-2023-golang/internal/database"
	"github.com/lib/pq"
)

type DBRepository struct{}

var dbRepository *DBRepository

func NewDBRepository() *DBRepository {
	if dbRepository == nil {
		dbRepository = &DBRepository{}
	}

	return dbRepository
}

func (m *DBRepository) FindById(id int) (p Person, err error) {
	sql := `SELECT id, apelido, nome, nascimento, stack
		FROM pessoas
		WHERE id = $1`

	var nascimento time.Time

	err = database.DB.QueryRow(sql, id).Scan(&p.Id, &p.Apelido, &p.Nome, &nascimento, pq.Array(&p.Stack))

	p.Nascimento = customDate{Time: nascimento}

	return p, err
}

func (m *DBRepository) FindByTerm(t string) (persons []Person, err error) {
	sql := `SELECT id, apelido, nome, nascimento, stack
		FROM pessoas
		WHERE search_index ILIKE '%'||$1||'%'`

	rows, err := database.DB.Query(sql, t)
	if err != nil {
		return persons, err
	}

	defer rows.Close()

	for rows.Next() {
		var nascimento time.Time
		var p Person

		err = rows.Scan(&p.Id, &p.Apelido, &p.Nome, &nascimento, pq.Array(&p.Stack))
		if err != nil {
			return persons, err
		}

		p.Nascimento = customDate{Time: nascimento}

		persons = append(persons, p)
	}

	return persons, nil
}

func (m *DBRepository) Count() (count int, err error) {
	sql := `SELECT COUNT(*) FROM pessoas`

	err = database.DB.QueryRow(sql).Scan(&count)

	return count, err
}

func (m *DBRepository) Insert(p Person) (id int, err error) {
	sql := `INSERT INTO pessoas (apelido, nome, nascimento, stack, search_index)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	searchIndex := strings.ToLower(fmt.Sprintf("%s %s %s", p.Apelido, p.Nome, strings.Join(p.Stack, " ")))

	err = database.DB.QueryRow(sql, p.Apelido, p.Nome, p.Nascimento.Time, pq.Array(p.Stack), searchIndex).Scan(&id)

	return id, err
}
