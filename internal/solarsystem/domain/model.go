package domain

type Body struct {
	Id          int    `json:"id" DB:"id"`
	Type        string `json:"type" DB:"type"`
	Name        string `json:"name" DB:"name"`
	Description string `json:"description" DB:"description"`
	Moons       int    `json:"moons" DB:"moons"`
}

type PhysicalCharacteristic struct {
	Id             int     `json:"id" DB:"id"`
	BodyId         int     `json:"body_id" DB:"body_id"`
	Density        float32 `json:"density" DB:"density"`
	Gravity        float32 `json:"gravity" DB:"gravity"`
	MassValue      float32 `json:"mass_value" DB:"mass_value"`
	MassExponent   int     `json:"mass_exponent" DB:"mass_exponent"`
	VolumeValue    float32 `json:"volume_value" DB:"volume_value"`
	VolumeExponent float32 `json:"volume_exponent" DB:"volume_exponent"`
}

type OrbitalParameters struct {
	Id              int     `json:"id" DB:"id"`
	BodyId          int     `json:"body_id" DB:"body_id"`
	SideralOrbit    float32 `json:"sideral_orbit" DB:"sideral_orbit"`
	SideralRotation float32 `json:"sideral_rotation" DB:"sideral_rotation"`
}
