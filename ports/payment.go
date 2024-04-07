package ports

import (
	"context"

	"github.com/ecommerce/adapters/secondary/postgres"
	"github.com/ecommerce/core/domain/payment"
)

type PaymentRepository interface {
	AddPayment(ctx context.Context, p payment.Payment) (*payment.Payment, error)
	GetById(ctx context.Context, id string) (*payment.Payment, error)
	GetAllPaymentsByUser(context context.Context, userId string, params postgres.QueryParams) (*[]payment.Payment, *int64, error)
	DeleteById(ctx context.Context, id string, p payment.Payment) (*payment.Payment, error)
}
