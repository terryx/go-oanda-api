package oanda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi_PlaceOrder(t *testing.T) {
	api := StubResponse("../fixture/order-created.json")
	req := OrderForm{}
	req.Order.Type = "LIMIT"
	req.Order.Instrument = "NAS100_USD"
	req.Order.Price = "5000.00"
	req.Order.Units = 1
	req.Order.TimeInForce = "GTC"
	req.Order.PositionFill = "DEFAULT"
	req.Order.StopLossOnFill.Price = "4000.00"
	req.Order.TakeProfitOnFill.Price = "9000.00"

	result, _ := api.PlaceOrder(req)
	assert.Equal(t, result.OrderCreateTransaction.Instrument, "NAS100_USD")
}

func TestApi_CancelOrder(t *testing.T) {
	api := StubResponse("../fixture/order-canceled.json")
	result, _ := api.CancelOrder("32")

	assert.Equal(t, result.OrderCancelTransaction.ID, "32")
}
