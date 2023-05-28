package body

type ServiceImpl struct {
	repository Repository
}

type Service interface {
	Get(id int) (*Body, error)
	Create(body Body) (*Body, error)
}

func NewService(repository Repository) *ServiceImpl {
	return &ServiceImpl{repository: repository}
}

func (s *ServiceImpl) Get(id int) (*Body, error) {
	return s.repository.Get(id)
}

func (s *ServiceImpl) Create(body Body) (*Body, error) {
	return s.repository.Create(body)
}
