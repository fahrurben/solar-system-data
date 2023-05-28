package physicalcharateristic

type PhysicalCharacteristic struct {
	Id             int     `json:"id" db:"id"`
	BodyId         int     `json:"body_id" db:"body_id"`
	Density        float32 `json:"density" db:"density"`
	Gravity        float32 `json:"gravity" db:"gravity"`
	MassValue      float32 `json:"mass_value" db:"mass_value"`
	MassExponent   int     `json:"mass_exponent" db:"mass_exponent"`
	VolumeValue    float32 `json:"volume_value" db:"volume_value"`
	VolumeExponent float32 `json:"volume_exponent" db:"volume_exponent"`
}
