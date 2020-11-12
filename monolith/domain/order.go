package domain

import (
	"context"
	"errors"
	"time"
)

type OrderID struct {
	ID string
}

type PayStatus int

var ErrInvalidServerPayStatus = errors.New("invalid server pay status")
var ErrOrderPayInvalid = errors.New("pay order invalid")

const (
	MissInitPayStatus PayStatus = iota
	UnPayStatus
	HasPayStatus
)

func (ps PayStatus) Validate() error {
	if ps != UnPayStatus || ps != HasPayStatus {
		return ErrInvalidServerPayStatus
	}
	return nil
}

func (ps PayStatus) String() string {
	switch ps {
	case MissInitPayStatus:
		return "pay status miss init"
	case UnPayStatus:
		return "unpay status"
	case HasPayStatus:
		return "haspay status"
	}
	return ""
}

type Order struct {
	ID        OrderID
	CreateAt  time.Time
	PayAt     time.Time
	PayStatus PayStatus
	ProductID ProductID
}

func NewOrder(product ProductID) Order {
	return Order{
		CreateAt:  time.Now(),
		PayStatus: UnPayStatus,
		ProductID: product,
	}
}

func (order Order) IsWaitForPay() bool {
	return order.PayStatus == UnPayStatus
}

func (order Order) OrderHasPay() bool {
	return order.PayStatus == HasPayStatus
}

func (order *Order) PayOrder() error {
	if order.IsWaitForPay() {
		order.PayAt = time.Now()
		order.PayStatus = HasPayStatus
	}
	return ErrOrderPayInvalid
}

type OrderRepo interface {
	GetOrder(OrderID) (Order, error)
	AllOrder() ([]Order, error)
	Store(*Order) error
	UpdateOrder(
		ctx context.Context,
		updateOrderID OrderID,
		updateFn func(*Order) (*Order, error),
	) error
}
