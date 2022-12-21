package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "stock"
)

func main() {

	router := gin.Default()
	router.GET("/products", listProducts)
	router.POST("/addProducts", createProduct)
	router.GET("/productsById", listProductsById)
	router.Run("localhost:9090")
}
