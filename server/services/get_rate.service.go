package services

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func (_this *CurrencyService) GetRate(base string, destination string) (*Rate, error) {
	_this.base = base
	serverAddr := _this.ConversionAddress(base, destination)
	_this.logger.Info("Calling Data Service:", zap.String("Server: ", serverAddr))
	resp, err := http.DefaultClient.Get(serverAddr)
	if err != nil {
		_this.logger.Error("An error occurred", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_this.logger.Error("An error occurred", zap.Int("Status Code: ", resp.StatusCode))
		return nil, err
	}
	response, err := _this.parseCurrencyValue(resp.Body, destination)
	if err != nil {
		return nil, err
	}
	dateStandard, err := time.Parse("2006-01-02", response.date)
	if err != nil {
		return nil, err
	}

	_this.time = dateStandard
	rate := (*response.rates)[destination]
	if rate == 0.0 {
		return nil, fmt.Errorf("invalid rate %f", rate)
	}

	return &Rate{Rate: rate, Destination: destination, Date: _this.GetDate(), Inverse: 1 / rate}, nil
}
