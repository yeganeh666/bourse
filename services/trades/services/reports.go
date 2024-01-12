package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	api "mabna/services/trades/api/src"
	"mabna/services/trades/repository"
)

type ReportsService interface {
	RegisterService(server *grpc.Server)
	GetLatestTrades(ctx context.Context, req *api.EmptyRequest) (*api.TradesReport, error)
}
type ReportServiceImpl struct {
	log               *log.Logger
	ReportsRepository repository.ReportsRepository
}

func NewReportService(log *log.Logger,
	reportsRepository repository.ReportsRepository) ReportsService {
	return &ReportServiceImpl{
		log:               log,
		ReportsRepository: reportsRepository,
	}
}

func (s *ReportServiceImpl) RegisterService(server *grpc.Server) {
	api.RegisterReportServiceServer(server, s)
}

func (s *ReportServiceImpl) GetLatestTrades(ctx context.Context, _ *api.EmptyRequest) (*api.TradesReport, error) {
	latestTrades, err := s.ReportsRepository.GetLatestTrades(ctx)
	if err != nil {
		s.log.WithError(err).Error("failed to get latest trades")
		return nil, err
	}
	response := new(api.TradesReport)
	for _, trade := range latestTrades {
		response.LatestTrades = append(response.LatestTrades, &api.LatestTrade{
			TradeId:        int32(trade.TradeID),
			InstrumentName: trade.InstrumentName,
			DateEn:         trade.DateEn,
			Open:           trade.Open,
			High:           trade.High,
			Low:            trade.Low,
			Close:          trade.Close,
		})
	}
	return response, nil
}
