package repository

import "github.com/meshenka/gokur/pkg/model"

// BusinessStore interface for the store
type BusinessStore interface {
	Init() error
	GetByID() *model.Business
}
