package handler

import (
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
)

func (_this *ServerHandler) StartPolling() {
	for {
		select {
		case <-_this.ticker.C:
			// When the ticker fires, it's time to harvest
			// loop oversubscribed clients
			for k, v := range *_this.subscriptions {
				// loop oversubscribed rates
				for _, rr := range *v {
					r, err := _this.ListRates(k.Context(), &rr)
					if err != nil {
						_this.logger.Error("Unable to get update rate", zap.String("base", rr.GetBase()), zap.Error(err))
					}
					err = k.Send(r)
					if err != nil {
						_this.logger.Error("Unable to get update rate", zap.String("base", rr.GetBase()), zap.Error(err))
					}
				}
			}
			//case u := <-_this.add:
			//	// At any time (other than when we're harvesting),
			//	// we can process a request to add a new URL
			//	_this.urls = append(_this.urls, u)
		}
	}
}

func (h *PollRates) AddRates(rates *exchange_rate.ListRatesResponse) {
	// Adding a new URL is as simple as tossing it onto a channel.
	h.add <- rates
}
