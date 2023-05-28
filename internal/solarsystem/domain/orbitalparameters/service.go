package orbitalparameters

type ServiceImpl struct {
	repository Repository
}

type Service interface {
	Get(id int) (*OrbitalParameters, error)
	Create(body OrbitalParameters) (*OrbitalParameters, error)
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Get(id int) (*OrbitalParameters, error) {
	return s.repository.Get(id)
}

func (s *ServiceImpl) Create(body OrbitalParameters) (*OrbitalParameters, error) {
	return s.repository.Create(body)
}
