package postgres

import (
	"context"
	"errors"

	"github.com/ecommerce/core/domain/payment"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrorPaymentNotFound = errors.New("payment not found")
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(dbConn *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: dbConn,
	}
}

func (pa *PaymentRepository) GetAllPaymentsByUser(ctx context.Context,
	userId string, params QueryParams) (*[]payment.Payment, *int64, error) {
	var list []payment.Payment

	res := pa.db.Where("user_id = ?", userId).Limit(params.Limit).Offset(params.Offset).Find(&list)
	if res.Error != nil {
		return nil, nil, res.Error
	}

	return &list, &res.RowsAffected, nil
}

func (pa *PaymentRepository) GetById(ctx context.Context, id string) (*payment.Payment, error) {
	var p payment.Payment
	res := pa.db.Where("id = ?", id).First(&p)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, ErrorPaymentNotFound
	}

	return &p, nil
}

func (pa *PaymentRepository) AddPayment(ctx context.Context, p payment.Payment) (*payment.Payment, error) {
	res := pa.db.Create(&p)
	if res.Error != nil {
		return nil, res.Error
	}

	return &p, nil
}

func (pa *PaymentRepository) DeleteById(ctx context.Context, id string, p payment.Payment) (*payment.Payment, error) {
	res := pa.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&p)
	if res.Error != nil {
		return nil, res.Error
	}

	return &p, nil
}
