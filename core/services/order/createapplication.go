package orders

import (
	"context"
	"strconv"
	"time"

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

func (o *OrderService) GetById(ctx context.Context, id string) (*order.Order, error) {
	od, err := o.orderRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return od, nil
}

func (o *OrderService) GetOrderCount(ctx context.Context) (*int64, *int64, error) {
	co, err := o.orderRepository.GetOrderCount(ctx)
	if err != nil {
		return nil, nil, err
	}

	coP, errP := o.orderRepository.GetPaidOrderCount(ctx)
	if errP != nil {
		return nil, nil, errP
	}

	return co, coP, nil
}

func (o *OrderService) GetOrdersByPeriod(ctx context.Context) (*[]OrderMonthQuantity, error) {
	var toOr []OrderMonthQuantity

	inDate, enDate, mot := o.GetInitialAndFinalDates()
	ord, errO := o.orderRepository.GetOrdersByPeriod(ctx, inDate, enDate)
	if errO != nil {
		return nil, errO
	}
	if len(*ord) == 0 {
		return &toOr, nil
	}

	for i := 0; i < len(mot); i++ {
		var aux OrderMonthQuantity
		var cot int64

		aux.Month = mot[i]

		for j := 0; j < len(*ord); j++ {
			auxOr := *ord
			if auxOr[j].CreatedAt.Month().String() == mot[i] {
				cot++
			}
		}

		aux.Quantity = cot
		toOr = append(toOr, aux)
	}

	return &toOr, nil
}

func (o *OrderService) GetProfitByOrdersByMonths(ctx context.Context) (*[]MonthData, error) {
	var foTo []OrderTotalMonth
	var auxTo []MonthData

	inDate, enDate, mot := o.GetInitialAndFinalDates()
	ord, errO := o.orderRepository.GetOrdersByPeriod(ctx, inDate, enDate)
	if errO != nil {
		return nil, errO
	}
	if len(*ord) == 0 {
		return &auxTo, nil
	}

	for i := 0; i < len(*ord); i++ {
		auxOr := *ord
		foTo = append(foTo, OrderTotalMonth{
			OrderId:  auxOr[i].ID,
			Subtotal: auxOr[i].TotalValue * float64(auxOr[i].Qtd),
			Month:    auxOr[i].CreatedAt.Month().String(),
		})
	}

	for i := 0; i < len(mot); i++ {
		var auxM MonthData
		auxM.Month = mot[i]

		for j := 0; j < len(foTo); j++ {
			if foTo[j].Month == mot[i] {
				auxM.Data = append(auxM.Data, foTo[j])
			}
		}

		auxTo = append(auxTo, auxM)
	}

	return &auxTo, nil
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

func (o *OrderService) UpdatePayment(ctx context.Context,
	id string, paid string) (*order.Order, error) {
	or, errO := o.orderRepository.GetById(ctx, id)
	if errO != nil {
		return nil, errO
	}

	nPa, errP := strconv.ParseBool(paid)
	if errP != nil {
		return nil, errP
	}

	up, err := o.orderRepository.UpdateOrderPayment(ctx, *or, nPa)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func (o *OrderService) UpdateRate(ctx context.Context,
	id string, rate string) (*order.Order, error) {
	or, errO := o.orderRepository.GetById(ctx, id)
	if errO != nil {
		return nil, errO
	}

	nRa, errR := strconv.Atoi(rate)
	if errR != nil {
		return nil, errR
	}

	rt, errR := order.NewRate(nRa)
	if errR != nil {
		return nil, errR
	}

	up, err := o.orderRepository.UpdateOrderRate(ctx, *or, *rt)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func (o *OrderService) UpdateStatus(ctx context.Context,
	id string, status string) (*order.Order, error) {
	or, errO := o.orderRepository.GetById(ctx, id)
	if errO != nil {
		return nil, errO
	}

	st, errS := order.NewStatus(status)
	if errS != nil {
		return nil, errS
	}

	up, err := o.orderRepository.UpdateOrderStatus(ctx, *or, *st)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func (o *OrderService) GetInitialAndFinalDates() (time.Time, time.Time, []string) {
	inDate := time.Now()
	enDate := inDate.AddDate(0, -5, 0)
	mot := make([]string, 5)

	for i := 0; i < 5; i++ {
		aux := inDate.AddDate(0, -i, 0)
		mot[i] = aux.Month().String()
	}

	return inDate, enDate, mot
}
