package sample

import (
	"errors"

	"github.com/hokauz/go-clean-api/core/entity"
)

// Service interface
type Service struct {
	repo     Repository
	Messages map[string]string
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
		Messages: map[string]string{
			"example": "Exemplo de erro",
		},
	}
}

// ReadOne -
func (s *Service) ReadOne(id string) (data *entity.Sample, msg string, err error) {
	data, err = s.repo.ReadOne(id)
	if err != nil {
		// TODO add code e erro para esta ação vide mapa
		msg = ""
		return
	}

	// Injection Case
	if data.ID.Hex() != id {
		// TODO add code e erro para esta ação vide mapa
		err = errors.New("Inalid request")
		msg = "Invalid request"
		return
	}

	return
}

// ReadAll -
func (s *Service) ReadAll() (data []*entity.Sample, msg string, err error) {
	data, err = s.repo.ReadAll()
	if err != nil {
		// TODO add code e erro para esta ação vide mapa
		msg = ""
		return
	}
	return
}

// Create -
func (s *Service) Create(data *entity.Sample) (d *entity.Sample, msg string, err error) {
	id, err := s.repo.Create(data)
	if err != nil {
		// TODO add code e erro para esta ação vide mapa
		msg = ""
		return
	}

	d, err = s.repo.ReadOne(id)
	return
}

// Update -
func (s *Service) Update(id string, data *entity.Sample) (d *entity.Sample, msg string, err error) {
	d, err = s.repo.Update(id, data)
	if err != nil {
		// TODO add code e erro para esta ação vide mapa
		msg = ""
		return
	}
	return
}

// Delete -
func (s *Service) Delete(id string) (msg string, err error) {
	err = s.repo.Delete(id)
	if err != nil {
		// TODO add code e erro para esta ação vide mapa
		msg = ""
		return
	}
	return
}
