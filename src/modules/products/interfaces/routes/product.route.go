package routes

import (
	"fmt"

	"hexagonal-architecture-golang/src/bootstrap/database"
	"hexagonal-architecture-golang/src/modules/products/application"
	"hexagonal-architecture-golang/src/modules/products/infrastructure/mongodb"
	"hexagonal-architecture-golang/src/modules/products/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRoute struct {
}

func NewProductRoute() *ProductRoute {
	return &ProductRoute{}
}

func LoadRoutes(r *gin.RouterGroup) {
	fmt.Println("llego a routes")

	con := database.GetMongoConnection()
	// productInfrastructure, err := mysql.NewProductRepositoryImplMysql(con.Conn)
	productInfrastructure, err := mongodb.NewProductRepositoryImpl(con.Conn)
	if err != nil {
		panic(err)
	}

	productApplication := application.NewProductApplication(productInfrastructure)

	productController := controllers.NewProductController(productApplication)

	r.GET("/all", productController.FindAll)
	r.POST("/save", productController.Save)

}
