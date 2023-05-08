package domain

type ServiceImpl struct {
	repository Repository
}

func (s *ServiceImpl) createBody(body Body) (*Body, error) {
	return s.repository.CreateBody(body)
}

func (s *ServiceImpl) CreatePhysicalCharacteristic(characteristic PhysicalCharacteristic) (*PhysicalCharacteristic, error) {
	return s.repository.CreatePhysicalCharacteristic(characteristic)
}

func (s *ServiceImpl) CreateOrbitalParameters(parameters OrbitalParameters) (*OrbitalParameters, error) {
	return s.repository.CreateOrbitalParameter(parameters)
}
