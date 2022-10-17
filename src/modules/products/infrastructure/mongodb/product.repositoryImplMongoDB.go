package mongodb

import (
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
	"hexagonal-architecture-golang/src/modules/products/domain/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImplMongoDB struct {
	DB *mongo.Database
}

func NewProductRepositoryImpl() (repositories.ProductRepository, error) {
	return &ProductRepositoryImplMongoDB{}, nil
}

func (p *ProductRepositoryImplMongoDB) FindAll() ([]entities.Product, error) {
	return []entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) FindById(id int64) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) Save(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) Update(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) Delete(product entities.Product) error {
	return nil
}
