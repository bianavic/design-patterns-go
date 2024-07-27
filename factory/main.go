package main

import (
	"fmt"
	"myapp/products"
)

func main() {

	// no factory
	//product := products.Product{
	//	ProductName: "Smartphone",
	//	CreatedAt:   time.Now(),
	//	UpdatedAt:   time.Now(),
	//}
	//fmt.Printf("Product %s was created at %s", product.ProductName, product.CreatedAt.UTC())

	factory := products.Product{}

	product := factory.New()

	fmt.Println("My product was created at", product.CreatedAt.UTC())
}
