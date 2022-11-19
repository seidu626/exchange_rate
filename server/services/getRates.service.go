package services

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// GetCurrencyRates Get records from a server
func (_this *CurrencyService) GetCurrencyRates(base string) (*map[string]float64, error) {
	_this.base = base
	serverAddr := _this.ServerAddress(base)
	_this.logger.Info("Calling Data Service:", zap.String("Server: ", serverAddr))
	resp, err := http.DefaultClient.Get(serverAddr)
	if err != nil {
		_this.logger.Error("An error occurred", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_this.logger.Error("An error occurred", zap.Int("Status Code: ", resp.StatusCode))
		return nil, fmt.Errorf("an error occurred %d", resp.StatusCode)
	}
	response, err := _this.parseCurrencyResponse(resp.Body, _this.base)
	if err != nil {
		return nil, err
	}
	_this.rates = response.rates

	dateStandard, err := time.Parse("2006-01-02", response.date)
	if err != nil {
		return nil, err
	}

	_this.time = dateStandard
	return response.rates, nil
}
