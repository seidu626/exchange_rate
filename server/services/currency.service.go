package services

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"time"
)

type CurrencyService struct {
	logger         *zap.Logger
	time           time.Time
	rates          *map[string]float64
	base           string
	serverAddress  string
	listServerAddr string
	conversionAddr string
}

type Currency struct {
	Name        string
	Description string
	Rate        float64
}

type Rate struct {
	Destination string
	Rate        float64
	Inverse     float64
	Date        *time.Time
}

func (_this *Rate) ComputeInverseRate(baseRate float64) float64 {
	_this.Inverse = baseRate / _this.Rate
	return _this.Inverse
}

type CurrencyResponse struct {
	date  string
	rates *map[string]float64
}

type ListCurrencyRequest struct {
	Base      string
	PageIndex int32
	PageSize  int32
}

type ListCurrencyResponse struct {
	Base      string
	PageIndex int32
	PageSize  int32
	Total     int32
	Data      *[]Currency
}

func NewCurrencyService(logger *zap.Logger, base string) *CurrencyService {
	return &CurrencyService{logger: logger,
		base:           base,
		listServerAddr: "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies.min.json",
		rates:          &map[string]float64{}, time: time.Now()}
}

// SetBase set base currency
func (_this *CurrencyService) SetBase(base string) {
	_this.base = base
}

// GetDate return rates date
func (_this *CurrencyService) GetDate() *time.Time {
	return &_this.time
	//layout := "2006-01-02T15:04:05.000Z"
	//dateStr := strings.Split(_this.time, " ")
	//t, err := time.Parse(layout, fmt.Sprintf("%sT%sZ", dateStr[0], dateStr[1]))
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return &t
}

// GetRates return base rates
func (_this *CurrencyService) GetRates() *map[string]float64 {
	(*_this.rates)[_this.base] = 1
	return _this.rates
}

// SetListServerAddress set new address, default: https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies.min.json"
func (_this *CurrencyService) SetListServerAddress(addr string) {
	_this.listServerAddr = addr
}

func (_this *CurrencyService) ServerAddress(base string) string {
	_this.serverAddress = fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%s.min.json", base)
	return _this.serverAddress
}

func (_this *CurrencyService) ConversionAddress(base string, destination string) string {
	_this.conversionAddr = fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%s/%s.json", base, destination)
	return _this.conversionAddr
}

// ParseCurrencyResponse converts xml response to map
func (_this *CurrencyService) parseCurrencyResponse(reader io.Reader, currency string) (*CurrencyResponse, error) {
	dec := json.NewDecoder(reader)
	result := &CurrencyResponse{}
	rec := map[string]interface{}{}
	for {
		var info map[string]interface{}
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			_this.logger.Error("", zap.Error(err))
			return nil, err
		}
		result.date = info["date"].(string)
		rec = info[currency].(map[string]interface{})
	}
	fmtRecords := map[string]float64{}
	for k, v := range rec {
		fmtRecords[k] = v.(float64)
	}
	result.rates = &fmtRecords
	return result, nil
}

// ParseCurrencyValue converts xml response to map
func (_this *CurrencyService) parseCurrencyValue(reader io.Reader, currency string) (*CurrencyResponse, error) {
	dec := json.NewDecoder(reader)
	result := &CurrencyResponse{}
	rec := 0.0
	for {
		var info map[string]interface{}
		if err := dec.Decode(&info); err == io.EOF {
			break
		} else if err != nil {
			_this.logger.Error("", zap.Error(err))
			return nil, err
		}
		result.date = info["date"].(string)
		rec = info[currency].(float64)
	}
	result.rates = &map[string]float64{currency: rec}
	return result, nil
}
