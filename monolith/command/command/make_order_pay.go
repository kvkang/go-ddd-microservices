package command

import (
	"context"

	"github.com/kvkang/go-ddd-microservices/monolith/domain"
)

type MakeOrderPayHandler struct {
	orderRepo domain.OrderRepo
}

func NewMakeOrderPayHandler(orderRepo domain.OrderRepo) MakeOrderPayHandler {
	return MakeOrderPayHandler{orderRepo}
}

func (handelr MakeOrderPayHandler) Handle(ctx context.Context, updateOrderID domain.OrderID) error {
	return handler.orderRepo(
		ctx,
		updateOrderID,
		func(updateBefore *domain.Order) (*domain.Order, error) {
			err := updateBefore.PayOrder()
			return updateBefore, err
		})
}
