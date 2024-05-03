package payments

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/payment"
)

func (p *PaymentService) GetAllPaymentsByUser(ctx context.Context,
	userId string, limit int, offset int) (*[]payment.Payment, error) {
	pay, _, err := p.paymentRepository.GetAllPaymentsByUser(ctx, userId, postgres.QueryParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	return pay, nil
}

func (p *PaymentService) CreatePayment(ctx context.Context, userId string,
	orderId string, toVa string,
	paVa string, paMe string) (*payment.Payment, error) {
	nToV, errT := payment.NewTotalValue(toVa)
	if errT != nil {
		return nil, errT
	}

	npaV, errP := payment.NewPaidValue(paVa)
	if errP != nil {
		return nil, errP
	}

	nPaM, errM := payment.NewPaymentMethod(paMe)
	if errM != nil {
		return nil, errM
	}

	nPay, err := payment.NewPayment(payment.Payment{
		UserID:        userId,
		OrderID:       orderId,
		TotalValue:    *nToV,
		PaidValue:     *npaV,
		PaymentMethod: *nPaM,
	})
	if err != nil {
		return nil, err
	}

	newP, err := p.paymentRepository.AddPayment(ctx, nPay)
	if err != nil {
		return nil, err
	}

	return newP, nil
}

func (p *PaymentService) DeletePayment(ctx context.Context, id string) (*payment.Payment, error) {
	pay, errG := p.paymentRepository.GetById(ctx, id)
	if errG != nil {
		return nil, errG
	}

	pay, err := p.paymentRepository.DeleteById(ctx, id, *pay)
	if err != nil {
		return nil, err
	}

	return pay, nil
}
