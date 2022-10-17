package repositories

import (
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(id int64) (entities.Product, error)
	Save(product entities.Product) (entities.Product, error)
	Update(product entities.Product) (entities.Product, error)
	Delete(product entities.Product) error
}
