package body

type Body struct {
	Id          int    `json:"id" db:"id"`
	Type        string `json:"type" db:"type"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Moons       int    `json:"moons" db:"moons"`
}
