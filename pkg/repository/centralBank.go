package repository

import (
	"GoAhead/pkg/data"
	"GoAhead/pkg/model"
	"context"
	"time"
)

type centralBankRepo struct {
	data *data.Data
}

func NewCentralBankRepo(data *data.Data) CentralBankRepo {
	return &centralBankRepo{
		data: data,
	}
}

type Valute struct {
	ID        uint      `gorm:"primarykey"`
	Vcurs     float64   `gorm:"column:curs"`
	VchCode   string    `gorm:"column:vch_code"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Valute) TableName() string {
	return "valutes"
}

// если нужно для записи после гет запроса
func (r centralBankRepo) GetCursByValue(ctx context.Context, valutes []model.Valute, val string) (*model.Valute, error) {
	return nil, nil
}

func (r centralBankRepo) GetDailyRate(ctx context.Context, valutes []model.Valute) error {
	tx := r.data.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	valutesInfo := make([]*Valute, len(valutes))

	for i, val := range valutes {
		valute := &Valute{
			Vcurs:     val.Vcurs,
			VchCode:   val.VchCode,
			CreatedAt: time.Now().UTC(),
		}

		valutesInfo[i] = valute
	}

	result := tx.Model(&Valute{}).Create(&valutesInfo)
	if result.Error != nil {
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
