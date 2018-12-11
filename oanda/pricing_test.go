package oanda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi_GetPricing(t *testing.T) {
	api := StubResponse("../fixture/pricing.json")
	result, _ := api.GetPricing("NAS100_USD")

	assert.Equal(t, result.Prices[0].Instrument, "NAS100_USD")
	assert.Equal(t, result.Prices[0].CloseoutAsk, "7276.9")
	assert.Equal(t, result.Prices[0].CloseoutBid, "7274.9")
}
