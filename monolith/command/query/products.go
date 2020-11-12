package query

import (
	"context"
	"time"

	"github.com/kvkang/go-ddd-microservices/monolith/domain"
	"github.com/sirupsen/logrus"
)

type ProductsHandler struct {
	domainRepo domain.ProductRepo
}

func NewProductsHandler(repo domain.ProductRepo) ProductsHandler {
	return ProductsHandler{repo}
}

func (handler ProductsHandler) Handle(ctx context.Context) (products []Product, err error) {
	start := time.Now()
	defer func() {
		logrus.
			WithError(err).
			WithField("duration", time.Since(start)).
			Debug("ProductsHandler executed")
	}()

	var domainProducts []domain.Product
	if domainProducts, err = handler.domainRepo.AllProduct(); err == nil {
		for i := range domainProducts {
			products = append(products, NewProduct(domainProducts[i]))
		}
	}
	return
}
