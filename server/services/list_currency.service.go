package services

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

// ListCurrencies List available currencies
func (_this *CurrencyService) ListCurrencies(request *ListCurrencyRequest) (*ListCurrencyResponse, error) {
	_this.base = request.Base
	serverAddr := _this.listServerAddr
	_this.logger.Info("Calling List Currency Service:", zap.String("Server: ", serverAddr))
	resp, err := http.DefaultClient.Get(serverAddr)
	if err != nil {
		_this.logger.Error("An error occurred", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_this.logger.Error("An error occurred", zap.Int("Status Code: ", resp.StatusCode))
		return nil, fmt.Errorf("an error occurred with status code %d", resp.StatusCode)
	}
	response := new(map[string]string)
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return pageResponse(request, response), nil
}

func pageResponse(request *ListCurrencyRequest, data *map[string]string) *ListCurrencyResponse {
	count := int32(0)
	records := &[]Currency{}
	for key, val := range *data {
		if count < request.PageIndex {
			count++
			continue
		}
		if int32(len(*records)) == request.PageSize {
			break
		}
		*records = append(*records, Currency{Name: key, Description: val})
	}
	response := &ListCurrencyResponse{PageIndex: request.PageIndex, PageSize: request.PageSize, Total: int32(len(*data)), Data: records}
	return response
}
