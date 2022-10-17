package infrastructure

import (
	"fmt"
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
	"hexagonal-architecture-golang/src/modules/products/domain/repositories"
)

type ProductInfrastructure struct {
}

func NewProductInfrastructure() (repositories.ProductRepository, error) {

	return &ProductInfrastructure{}, nil
}

func (p *ProductInfrastructure) FindAll() ([]entities.Product, error) {
	fmt.Println("llego a infrastructure")
	return []entities.Product{}, nil
}

func (p *ProductInfrastructure) FindById(id int64) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductInfrastructure) Save(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductInfrastructure) Update(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductInfrastructure) Delete(product entities.Product) error {
	return nil
}
