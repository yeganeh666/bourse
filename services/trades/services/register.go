package services

import (
	log "github.com/sirupsen/logrus"
	models "mabna/services/trades/api/src"
	"mabna/services/trades/configs"
	"mabna/shared/grpcext"
)

type Service struct {
	log            *log.Logger
	Config         *configs.Configs
	ReportsService models.ReportServiceClient
}

func NewService(log *log.Logger, conn *grpcext.Connection) *Service {

	return &Service{
		log:            log,
		ReportsService: models.NewReportServiceClient(conn),
	}
}
