package domain

type ProductID struct {
	ID string
}

type ProductName struct {
	Name string
}

type ProductDescribe struct {
	Describe string
}

type Product struct {
	ID       ProductID
	Name     ProductName
	Describe ProductDescribe
}

func NewProduct(name ProductName, des ProductDescribe) Product {
	return Product{Name: name, Describe: des}
}

type ProductRepo interface {
	Store(*Product) error
	GetProcut(ProductID) (Product, error)
	AllProduct() ([]Product, error)
}
