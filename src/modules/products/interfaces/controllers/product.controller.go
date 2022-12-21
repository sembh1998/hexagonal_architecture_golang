package controllers

import (
	"hexagonal-architecture-golang/src/modules/products/application"
	"hexagonal-architecture-golang/src/modules/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productApplication *application.ProductApplication
}

func NewProductController(productApplication *application.ProductApplication) *ProductController {
	return &ProductController{
		productApplication: productApplication,
	}
}

func (p *ProductController) FindAll(c *gin.Context) {
	products, err := p.productApplication.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"products": products,
	})

}

func (p *ProductController) Save(c *gin.Context) {
	prod := entities.Product{}
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := p.productApplication.Save(prod)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"product": product,
	})
}
