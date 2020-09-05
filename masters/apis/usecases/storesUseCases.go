package usecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type StoresUseCases interface {
	GetStores() ([]*models.StoreModels, error)
	GetStoreByID(ID string) (*models.StoreModels, error)
	PostStore(d models.StoreModels) (*models.StoreModels, error)
	UpdateStore(ID string, data models.StoreModels) (*models.StoreModels, error)
	DeleteStore(ID string) (*models.StoreModels, error)

	GetStoresCategory() ([]*models.StoreCategory, error)
}