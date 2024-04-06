package orders

import (
	"context"

	"github.com/ecommerce/core/domain/order"
)

func (o *OrderService) GetByUserId(ctx context.Context,
	userId string,
	limit int,
	offset int) (*[]order.Order, error) {
	ods, err := o.orderRepository.GetByUserId(ctx, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	return ods, nil
}

func (o *OrderService) AddOrder(ctx context.Context,
	userId string, prodId string, qtd int, toV float64) (*order.Order, error) {
	prs, errP := order.NewProductId(prodId)
	if errP != nil {
		return nil, errP
	}

	uId, errU := order.NewUserId(userId)
	if errU != nil {
		return nil, errU
	}

	toV1, errT := order.NewTotalValue(toV)
	if errT != nil {
		return nil, errT
	}

	qtd1, errQ := order.NewQuantity(qtd)
	if errQ != nil {
		return nil, errQ
	}

	sts, errS := order.NewStatus("PENDENTE")
	if errS != nil {
		return nil, errS
	}

	rt, errR := order.NewRate(0)
	if errR != nil {
		return nil, errR
	}

	or, errN := order.NewOrder(order.Order{
		ProductID:  prs,
		Qtd:        *qtd1,
		TotalValue: *toV1,
		Status:     *sts,
		Userid:     uId,
		Rate:       *rt,
		Paid:       false,
	})
	if errN != nil {
		return nil, errN
	}

	newO, err := o.orderRepository.AddOrder(ctx, or)
	if err != nil {
		return nil, err
	}

	return newO, nil
}
