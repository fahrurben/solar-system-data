package orbitalparameters

type OrbitalParameters struct {
	Id              int     `json:"id" db:"id"`
	BodyId          int     `json:"body_id" db:"body_id"`
	SideralOrbit    float32 `json:"sideral_orbit" db:"sideral_orbit"`
	SideralRotation float32 `json:"sideral_rotation" db:"sideral_rotation"`
}
