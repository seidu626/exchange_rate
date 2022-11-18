package services

import (
	"fmt"
	"go.uber.org/zap"
	"testing"
)

func TestCurrencyService_ListCurrencies(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	service := NewCurrencyService(logger, "eur")
	resp, err := service.ListCurrencies(&ListCurrencyRequest{
		PageIndex: 0, PageSize: 5})
	if err != nil {
		logger.Error("An error occurred", zap.Error(err))
	}
	fmt.Printf("%#v", resp.Data)
}

func TestCurrencyService_GetCurrency(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	service := NewCurrencyService(logger, "eur")
	rates, err := service.GetCurrencyRates("eur")
	if err != nil {
		logger.Error("An error occurred", zap.Error(err))
	}
	fmt.Printf("%#v", service.GetDate())
	fmt.Printf("%#v", *rates)
}

func TestCurrencyService_ConvertTo(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	service := NewCurrencyService(logger, "eur")
	value, err := service.GetRate("usd", "ghs")
	if err != nil {
		logger.Error("An error occurred", zap.Error(err))
	}
	fmt.Printf("%#v", value)

}
