package server

import (
	"hexagonal-architecture-golang/src/bootstrap"
	productRoutes "hexagonal-architecture-golang/src/modules/products/interfaces/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer() (bootstrap.Bootstrap, error) {
	return &Server{}, nil
}

func (s *Server) Initialice() error {

	router := gin.Default()

	productGroup := router.Group("/products")
	productRoutes.LoadRoutes(productGroup)

	router.Run(":8080")
	return nil
}
