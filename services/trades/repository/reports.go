package repository

import (
	"context"
)

type ReportsRepository interface {
	GetLatestTrades(ctx context.Context) ([]*LatestTrades, error)
}

type ReportsRepositoryImpl struct {
	*Repository
}

func NewReportsRepository(Repository *Repository) ReportsRepository {
	return &ReportsRepositoryImpl{
		Repository: Repository,
	}
}

type LatestTrades struct {
	TradeID        int `json:"trade_id"`
	InstrumentName string
	DateEn         string
	Open           float64
	High           float64
	Low            float64
	Close          float64
}

func (r *ReportsRepositoryImpl) GetLatestTrades(ctx context.Context) ([]*LatestTrades, error) {
	var trades []*LatestTrades

	err := r.DB.Table("trade").
		Select(" trade.id as trade_id, instrument.name AS instrument_name, trade.dateEn AS date_en, trade.Open, trade.High, trade.Low, trade.Close").
		Joins("JOIN instrument ON trade.InstrumentId = instrument.id").
		Where("trade.id IN (SELECT DISTINCT ON (InstrumentId) ID FROM trade ORDER BY InstrumentId, dateEn DESC)").
		Scan(&trades).Error

	return trades, err
}
