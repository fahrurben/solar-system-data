package physicalcharateristic

type ServiceImpl struct {
	repository Repository
}

type Service interface {
	Get(id int) (*PhysicalCharacteristic, error)
	Create(body PhysicalCharacteristic) (*PhysicalCharacteristic, error)
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Get(id int) (*PhysicalCharacteristic, error) {
	return s.repository.Get(id)
}

func (s *ServiceImpl) Create(body PhysicalCharacteristic) (*PhysicalCharacteristic, error) {
	return s.repository.Create(body)
}
