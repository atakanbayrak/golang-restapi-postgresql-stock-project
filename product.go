package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var products = []Product{}

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
	Type  int    `json:"type"`
}

func createProduct(context *gin.Context) {
	var productByUser Product
	err := context.BindJSON(&productByUser)

	if err == nil && productByUser.Name != "" && productByUser.Stock != 0 && productByUser.Type != 0 {
		products = append(products, productByUser)
		addProducts(&productByUser)
		context.IndentedJSON(http.StatusCreated, gin.H{"message": "Product has been created", "product_name": productByUser.Name})
		return
	} else {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Product has not been created"})
		return
	}
}
func addProducts(productByUser *Product) {

	postgreInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgreInfo)

	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection with db.Ping")
	}

	fmt.Println("Successful connection to database")

	sqlStatement := `INSERT INTO products (product_name,product_stock,category_type) VALUES ($1,$2,$3) `
	insert, err := db.Exec(sqlStatement, productByUser.Name, productByUser.Stock, productByUser.Type)
	if err != nil {
		panic(err.Error())
	}

	rowsAffec, _ := insert.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error inserting row:", err)
		return
	}

	lastInserted, _ := insert.LastInsertId()
	rowsAffected, _ := insert.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)

	fmt.Println("Successful Connection to Database!")

}
func listProducts(context *gin.Context) {
	postgreInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgreInfo)

	if err != nil {
		fmt.Println("Error validating sql.Open arguments")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection with db.Ping")
	}

	fmt.Println("Successful connection to database")

	rows, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var productsNew Product
	var productsShow = []Product{}

	var title1 int
	var title2 string
	var title3 int
	var title4 int

	for rows.Next() {

		if err := rows.Scan(&title1, &title2, &title3, &title4); err != nil {
			log.Fatal(err)
		}
		fmt.Println(title1)
		productsNew.Id = title1
		productsNew.Name = title2
		productsNew.Stock = title3
		productsNew.Type = title4
		productsShow = append(productsShow, productsNew)
	}

	rows.Close()
	context.IndentedJSON(http.StatusOK, productsShow)

}
func listProductsById(context *gin.Context) {
	var productById Product
	err := context.BindJSON(&productById)
	for i := 0; i < len(products); i++ {
		if productById.Id == products[i].Id && err != nil {
			context.IndentedJSON(http.StatusOK, gin.H{"message:": "Products has been found"})
			return
		}
	}

}
