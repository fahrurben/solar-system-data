package physicalcharateristic

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Get(id int) (*PhysicalCharacteristic, error)
	Create(body PhysicalCharacteristic) (*PhysicalCharacteristic, error)
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Get(id int) (*PhysicalCharacteristic, error) {
	physicalChar := &PhysicalCharacteristic{}
	err := r.db.Get(physicalChar, "select id, body_id, density, gravity, mass_value, mass_exponent, volume_value, volume_exponent from physical_data where id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return physicalChar, err
}

func (r *RepositoryImpl) Create(characteristic PhysicalCharacteristic) (*PhysicalCharacteristic, error) {
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
