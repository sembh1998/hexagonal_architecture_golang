package mongodb

import (
	"context"
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
	"hexagonal-architecture-golang/src/modules/products/domain/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImplMongoDB struct {
	DB *mongo.Database
}

func NewProductRepositoryImpl(conn *mongo.Database) (repositories.ProductRepository, error) {
	return &ProductRepositoryImplMongoDB{
		DB: conn,
	}, nil
}

func (p *ProductRepositoryImplMongoDB) FindAll() ([]entities.Product, error) {
	coll := p.DB.Collection("products")
	response, err := coll.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	var products []entities.Product
	if err = response.All(context.Background(), &products); err != nil {
		panic(err)
	}
	return products, nil
}

func (p *ProductRepositoryImplMongoDB) FindById(id int64) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) Save(product entities.Product) (entities.Product, error) {
	coll := p.DB.Collection("products")
	_, err := coll.InsertOne(context.Background(), product)
	if err != nil {
		return entities.Product{}, err
	}
	return entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) Update(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImplMongoDB) Delete(product entities.Product) error {
	return nil
}
