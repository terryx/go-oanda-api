package oanda

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApi_GetInstruments(t *testing.T) {
	api := StubResponse("../fixture/instruments.json")
	res, _ := api.GetInstruments()

	assert.Equal(t, res.Instruments[0].Name, "AUD_HKD")
}

func TestApi_GetInstrumentCandles(t *testing.T) {
	api := StubResponse("../fixture/candles.json")
	param := CandlestickParam{"NAS100_USD", "M5"}
	res, _ := api.GetInstrumentCandles(param)

	assert.Equal(t, res.Instrument, "NAS100_USD")
	assert.Equal(t, res.Granularity, "M5")
	assert.NotEmpty(t, res.Candles)
}
