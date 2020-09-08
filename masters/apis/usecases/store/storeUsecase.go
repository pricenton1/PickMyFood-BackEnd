package storeusecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type StoreUsecase interface {
	AddStore(store *models.Store) error
	GetStoreByID(id string) (*models.Store, error)
	GetAllStore() ([]*models.Store, error)
	UpdateStore(id string, Store *models.Store) error
	DeleteStore(id string) error
	Auth(username string) (*models.Store, error)
}