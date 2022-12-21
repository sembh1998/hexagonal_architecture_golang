package application

import (
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
	"hexagonal-architecture-golang/src/modules/products/domain/repositories"
)

type ProductApplication struct {
	productRepository repositories.ProductRepository
}

func NewProductApplication(productRepository repositories.ProductRepository) *ProductApplication {
	return &ProductApplication{
		productRepository: productRepository,
	}
}

func (p *ProductApplication) FindAll() ([]entities.Product, error) {
	return p.productRepository.FindAll()
}

func (p *ProductApplication) Save(product entities.Product) (entities.Product, error) {
	return p.productRepository.Save(product)
}
