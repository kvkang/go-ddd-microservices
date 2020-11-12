package query

import "github.com/kvkang/go-ddd-microservices/monolith/domain"

type Product struct {
	ID       string
	Name     string
	Describe string
}

func NewProduct(p domain.Product) Product {
	return Product{}
}
