package service

import (
	"GoAhead/pkg/model"
	"GoAhead/pkg/repository"
	"context"
)

type CentralBankUseCase struct {
	repo repository.CentralBankRepo
}

func NewCentralBankUseCase(repo repository.CentralBankRepo) *CentralBankUseCase {
	return &CentralBankUseCase{repo: repo}
}

var currencyCodes = []string{
	"AUD", "AZN", "GBP", "AMD", "BYN", "BGN", "BRL", "HUF", "VND", "HKD",
	"GEL", "DKK", "AED", "USD", "EUR", "EGP", "INR", "IDR", "KZT", "CAD",
	"QAR", "KGS", "CNY", "MDL", "NZD", "NOK", "PLN", "RON", "XDR", "SGD",
	"TJS", "THB", "TRY", "TMT", "UZS", "UAH", "CZK", "SEK", "CHF", "RSD",
	"ZAR", "KRW", "JPY",
}

func (uc CentralBankUseCase) GetCursByValue(_ context.Context, valutes []model.Valute, val string) (*model.Valute, error) {
	var response *model.Valute

	for _, v := range valutes {
		if v.VchCode == val {
			response = &v
			break
		}
	}

	return response, nil
}

func (uc CentralBankUseCase) GetDailyRate(ctx context.Context, valutes []model.Valute) error {
	return uc.repo.GetDailyRate(ctx, valutes)
}
