package rinha

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Pessoa struct {
	Id         int        `json:"id"`
	Apelido    string     `json:"apelido"`
	Nome       string     `json:"nome"`
	Nascimento customDate `json:"nascimento"`
	Stack      []string   `json:"stack"`
}

type customDate struct {
	time.Time
}

func (t *customDate) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(`"2006-01-02"`, string(b))
	t.Time = time

	return err
}

func (t *customDate) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(`"2006-01-02"`)), nil
}

func InsertPerson(db *sql.DB, p *Pessoa) (id int, err error) {
	sql := `INSERT INTO pessoas (apelido, nome, nascimento, stack)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err = db.QueryRow(sql, p.Apelido, p.Nome, p.Nascimento.Time, pq.Array(p.Stack)).Scan(&id)

	return id, err
}

func GetPerson(db *sql.DB, id string) (Pessoa, error) {
	sql := `SELECT id, apelido, nome, nascimento, stack
		FROM pessoas
		WHERE id = $1`

	var nascimento time.Time

	var p Pessoa
	err := db.QueryRow(sql, id).Scan(&p.Id, &p.Apelido, &p.Nome, &nascimento, pq.Array(&p.Stack))

	p.Nascimento = customDate{Time: nascimento}

	return p, err
}
