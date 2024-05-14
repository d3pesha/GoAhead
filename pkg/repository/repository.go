package repository

import (
	"GoAhead/pkg/model"
	"context"
)

type CentralBankRepo interface {
	GetCursByValue(context.Context, []model.Valute, string) (*model.Valute, error)
	GetDailyRate(context.Context, []model.Valute) error
}
