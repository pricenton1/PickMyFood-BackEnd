package productUsecases

import "github.com/inact25/PickMyFood-BackEnd/masters/apis/models"

type ProductUsecase interface {
	AddProduct(storeID string, product *models.Product) error
	GetProductByID(id string) (*models.Product, error)
	GetAllProductByStore(storeID string) ([]*models.Product, error)
	UpdateProductWithPrice(id string, product *models.Product) error
	DeleteProduct(id string) error
}