package sample

import "github.com/hokauz/go-clean-api/core/entity"

type (
	// Reader -
	Reader interface {
		ReadOne(id string) (*entity.Sample, error)
	}

	// Writer -
	Writer interface {
		Create(data *entity.Sample) (string, error)
		Update(id string, data *entity.Sample) (*entity.Sample, error)
		Delete(id string) error
	}

	// UseCase -
	UseCase interface {
		Reader
		Writer
	}

	// Repository -
	Repository interface {
		Reader
		Writer
	}
)
