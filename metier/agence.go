package metier

type Agence struct {
	Code string
	Nom  string
}

type AgenceRepository interface {
	Get(code string) Agence
}

type Service struct {
	repository AgenceRepository
}

func NewService(repository AgenceRepository) Service {
	return Service{repository: repository}
}

func (s Service) GetAgence(code string) Agence {
	return s.repository.Get(code)
}
