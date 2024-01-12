package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	reportsAPI "mabna/services/trades/services"
	"mabna/shared/grpcext"
)

type ReportsHandlerImpl struct {
	log     *logrus.Logger
	service *reportsAPI.Service
}

func NewReportsHandler(log *logrus.Logger, reportsAddr string) *ReportsHandlerImpl {
	reportsConn := grpcext.NewConnection(reportsAddr)
	fmt.Println(reportsConn)
	return &ReportsHandlerImpl{
		log:     log,
		service: reportsAPI.NewService(log, reportsConn),
	}
}
