package command

import (
	"context"

	"github.com/kvkang/go-ddd-microservices/monolith/domain"
)

type CreateUnpayOrderHandler struct {
	productRepo domain.ProductRepo
	orderRepo   domain.OrderRepo
}

func NewCreateUnpayOrderHandler(productRepo domain.ProductRepo, orderRepo domain.OrderRepo) error {
	return CreateUnpayOrderHandler{productRepo, orderRepo}
}

func (handler CreateUnpayOrderHandler) Handle(ctx context.Context, productID domain.ProductID) error {
	if _, err := handler.productRepo.GetProcut(productID); err != nil {
		return err
	}

	order := domain.NewOrder(productID)
	return handler.orderRepo.Store(&order)
}
