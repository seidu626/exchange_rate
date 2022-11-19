package handler

import (
	"fmt"
	exchange_rate "github.com/seidu626/exchange_rate/proto"
	"go.uber.org/zap"
	"time"
)

func (_this *ServerHandler) Send(response *exchange_rate.ListRatesResponse) error {
	_this.logger.Info("Send subscriptions", zap.String("", fmt.Sprintf("%#v", response)))
	return nil
}

// Send checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes

func (_this *ServerHandler) SendUpdate(interval time.Duration) error {
	ret := make(chan struct{})
	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("Ticker after every %d", interval)
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				//for k, v := range e.rates {
				//	// change can be 10% of original value
				//	change := (rand.Float64() / 10)
				//	// is this a postive or negative change
				//	direction := rand.Intn(1)
				//
				//	if direction == 0 {
				//		// new value with be min 90% of old
				//		change = 1 - change
				//	} else {
				//		// new value will be 110% of old
				//		change = 1 + change
				//	}
				//
				//	// modify the rate
				//	e.rates[k] = v * change
				//}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()
	return nil
}
