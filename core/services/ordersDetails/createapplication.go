package ordersdetails

import (
	"context"

	"github.com/ecommerce/core/domain/orderDetails"
)

func (od *OrdersDetailsService) GetTotalOrderValueAndQuatity(ctx context.Context, det []orderDetails.OrderProductData) (*float64, *int, error) {
	if len(det) == 0 {
		return nil, nil, ErrorEmptyOrderDetailsList
	}

	var toV float64
	var qtd int
	for _, order := range det {
		toV += order.Value
		qtd += order.Quantity
	}

	return &toV, &qtd, nil
}

func (od *OrdersDetailsService) CheckOrderDetails(ctx context.Context, orderId string,
	det orderDetails.OrderProductData) (*orderDetails.OrderDetails, error) {
	oId, errO := orderDetails.NewOrderId(orderId)
	if errO != nil {
		return nil, errO
	}

	pId, errP := orderDetails.NewProductId(det.ProductId)
	if errP != nil {
		return nil, errP
	}

	val, errV := orderDetails.NewValue(det.Value)
	if errV != nil {
		return nil, errV
	}

	qtd, errQ := orderDetails.NewQuantity(det.Quantity)
	if errQ != nil {
		return nil, errQ
	}

	odn, errN := orderDetails.NewOrderDetails(orderDetails.OrderDetails{
		OrderID:   oId,
		Value:     *val,
		Qtd:       *qtd,
		ProductID: *pId,
	})
	if errN != nil {
		return nil, errN
	}

	return &odn, nil
}
