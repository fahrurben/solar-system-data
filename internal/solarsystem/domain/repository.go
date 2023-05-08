package domain

import "github.com/jmoiron/sqlx"

type Repository interface {
	CreateBody(body Body) (*Body, error)
	CreatePhysicalCharacteristic(characteristic PhysicalCharacteristic) (*PhysicalCharacteristic, error)
	CreateOrbitalParameter(parameter OrbitalParameters) (*OrbitalParameters, error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) CreateBody(body Body) (*Body, error) {
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

func (r *RepositoryImpl) CreatePhysicalCharacteristic(characteristic PhysicalCharacteristic) (*PhysicalCharacteristic, error) {
	_, err := r.db.Exec(
		"insert into physical_data (id, body_id, density, gravity, mass_value, mass_exponent, volume_value, volume_exponent) values (?, ?, ?, ?, ?, ?, ?, ?)",
		characteristic.Id,
		characteristic.BodyId,
		characteristic.Density,
		characteristic.Gravity,
		characteristic.MassValue,
		characteristic.MassExponent,
		characteristic.VolumeValue,
		characteristic.VolumeExponent,
	)

	if err != nil {
		return nil, err
	}

	return &characteristic, err
}

func (r *RepositoryImpl) CreateOrbitalParameter(parameter OrbitalParameters) (*OrbitalParameters, error) {
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
