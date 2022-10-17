package routes

import (
	"fmt"

	"hexagonal-architecture-golang/src/bootstrap/database"
	"hexagonal-architecture-golang/src/modules/products/application"
	"hexagonal-architecture-golang/src/modules/products/infrastructure/mysql"
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

	con := database.GetMysqlConnection()
	productInfrastructure, err := mysql.NewProductRepositoryImplMysql(con.Conn)
	if err != nil {
		panic(err)
	}
	// productMongoDB, err := mongodb.NewProductRepositoryImpl()
	// if err != nil {
	// 	panic(err)
	// }

	productApplication := application.NewProductApplication(productInfrastructure)

	productController := controllers.NewProductController(productApplication)

	r.GET("/employees", productController.FindAll)

}
