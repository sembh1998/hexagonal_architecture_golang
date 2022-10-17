package controllers

import (
	"fmt"
	"hexagonal-architecture-golang/src/modules/products/application"

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
	fmt.Println("llego a controller")
	products, err := p.productApplication.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"products": products,
	})

}
