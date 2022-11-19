package subscriptionClient

import exchange_rate "github.com/seidu626/exchange_rate/proto"

func (_this *Handler) Send(request *exchange_rate.ListRatesRequest) error {
	err := _this.ExchangeRates_SubscriptionClient.Send(request)
	if err != nil {
		return err
	}
	return nil
}
