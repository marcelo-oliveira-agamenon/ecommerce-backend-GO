package payments

import (
	"context"

	"github.com/ecommerce/core/domain/payment"
	"github.com/ecommerce/ports"
)

type API interface {
	GetAllPaymentsByUser(ctx context.Context, userId string, limit int, offset int) (*[]payment.Payment, error)
	CreatePayment(ctx context.Context, userId string, orderId string, toVa string, paVa string, paMe string) (*payment.Payment, error)
	DeletePayment(ctx context.Context, id string) (*payment.Payment, error)
}

type PaymentService struct {
	paymentRepository ports.PaymentRepository
}

func NewPaymentService(pay ports.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepository: pay,
	}
}
