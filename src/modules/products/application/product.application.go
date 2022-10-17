package application

import (
	"fmt"
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
	fmt.Println("llego a application")
	return p.productRepository.FindAll()
}
