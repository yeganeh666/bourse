package handlers

import (
	reportsAPI "mabna/services/trades/api/src"

	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleLatestTrades
// @Summary HandleLatestTrades
// @Description retrieves the latest trades report for financial instruments
// @Tags Objects
// @Accept json
// @Produce json
// @Success 200 {object} reportsAPI.TradesReport
// @Router /trades/latest-report [get]
func (h *ReportsHandlerImpl) HandleLatestTrades(ctx *gin.Context) {
	objectDetails, err := h.service.ReportsService.GetLatestTrades(ctx, &reportsAPI.EmptyRequest{})
	if err != nil {
		h.log.WithError(err).Error("error while get latest trades")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, objectDetails)
}
