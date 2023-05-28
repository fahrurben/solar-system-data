package orbitalparameters

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Get(id int) (*OrbitalParameters, error)
	Create(body OrbitalParameters) (*OrbitalParameters, error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Get(id int) (*OrbitalParameters, error) {
	orbitalParameter := &OrbitalParameters{}
	err := r.db.Get(orbitalParameter, "select id, body_id, sideral_orbit, sideral_rotation from orbital_parameters where id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return orbitalParameter, err
}

func (r *RepositoryImpl) Create(parameter OrbitalParameters) (*OrbitalParameters, error) {
	_, err := r.db.Exec(
		"insert into orbital_parameters (id, body_id, sideral_orbit, sideral_rotation) values (?, ?, ?, ?)",
		parameter.Id,
		parameter.BodyId,
		parameter.SideralOrbit,
		parameter.SideralRotation,
	)

	if err != nil {
		return nil, err
	}

	return &parameter, err
}
