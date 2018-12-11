package oanda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi_CloseTrade(t *testing.T) {
api := StubResponse("../fixture/trade-closed.json")
result, _ := api.CloseTrade("4119")

assert.Equal(t, result.OrderFillTransaction.TradesClosed[0].TradeID, "4119")
}