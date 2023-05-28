package body

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Get(id int) (*Body, error)
	Create(body Body) (*Body, error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Get(id int) (*Body, error) {
	body := &Body{}
	err := r.db.Get(body, "select id, type, name, description, moons from body where id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return body, err
}

func (r *RepositoryImpl) Create(body Body) (*Body, error) {
	_, err := r.db.Exec(
		"insert into body (id, type, name, description, moons) values (?, ?, ?, ?, ?)",
		body.Id,
		body.Type,
		body.Name,
		body.Description,
		body.Moons,
	)

	if err != nil {
		return nil, err
	}

	return &body, err
}
