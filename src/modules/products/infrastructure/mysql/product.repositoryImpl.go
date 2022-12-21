package mysql

import (
	"context"
	"database/sql"
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
	"hexagonal-architecture-golang/src/modules/products/domain/repositories"
	productsqlc "hexagonal-architecture-golang/src/modules/products/infrastructure/mysql/sqlc"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepositoryImplMysql(DB *sql.DB) (repositories.ProductRepository, error) {
	return &ProductRepositoryImpl{
		DB: DB,
	}, nil
}

func (p *ProductRepositoryImpl) FindAll() ([]entities.Product, error) {
	db := productsqlc.New(p.DB)

	products, err := db.FindAll(context.Background())
	if err != nil {
		return nil, err
	}
	entProducts := []entities.Product{}

	for _, product := range products {
		var price *float64
		if product.Price.Valid {
			price = &product.Price.Float64
		} else {
			price = nil
		}
		id64 := int64(product.ID)
		entProducts = append(entProducts, entities.Product{
			ID:          &id64,
			Name:        product.Name,
			Description: product.Description,
			Price:       price,
		})
	}

	return entProducts, nil
}

func (p *ProductRepositoryImpl) FindById(id int64) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImpl) Save(product entities.Product) (entities.Product, error) {
	db := productsqlc.New(p.DB)

	result, err := db.Save(context.Background(), productsqlc.SaveParams{
		Name:        product.Name,
		Description: product.Description,
		Price: sql.NullFloat64{
			Float64: *product.Price,
			Valid:   *product.Price != 0,
		},
	})
	if err != nil {
		return entities.Product{}, err
	}
	last_id, err := result.LastInsertId()
	if err != nil {
		return entities.Product{}, err
	}
	product.ID = &last_id
	return product, nil
}

func (p *ProductRepositoryImpl) Update(product entities.Product) (entities.Product, error) {
	return entities.Product{}, nil
}

func (p *ProductRepositoryImpl) Delete(product entities.Product) error {
	return nil
}
