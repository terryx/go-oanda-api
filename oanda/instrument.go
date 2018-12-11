package oanda

import (
	"fmt"
	"net/url"
)

type Instrument struct {
	Instruments []struct {
		Name                        string  `json:"name"`
		Type                        string  `json:"type"`
		DisplayName                 string  `json:"displayName"`
		PipLocation                 float64 `json:"pipLocation"`
		DisplayPrecision            float64 `json:"displayPrecision"`
		TradeUnitsPrecision         float64 `json:"tradeUnitsPrecision"`
		MinimumTradeSize            string  `json:"minimumTradeSize"`
		MaximumTrailingStopDistance string  `json:"maximumTrailingStopDistance"`
		MinimumTrailingStopDistance string  `json:"minimumTrailingStopDistance"`
		MaximumPositionSize         string  `json:"maximumPositionSize"`
		MaximumOrderUnits           string  `json:"maximumOrderUnits"`
		MarginRate                  string  `json:"marginRate"`
	} `json:"instruments"`
}

type Chart struct {
	Instrument  string `json:"instrument"`
	Granularity string `json:"granularity"`
	Candles     []struct {
		Complete bool   `json:"complete"`
		Volume   int    `json:"volume"`
		Time     string `json:"time"`
		Mid      struct {
			Open  float64 `json:"o"`
			High  float64 `json:"h"`
			Low   float64 `json:"l"`
			Close float64 `json:"c"`
		} `json:"mid"`
	} `json:"candles"`
}

type CandlestickParam struct {
	Name        string
	Granularity string
}

func (a *Api) GetInstrumentCandles(param CandlestickParam) (chart *Chart, err error) {
	req, _ := a.NewRequest("GET", fmt.Sprintf("v3/instruments/%s/candles", param.Name), nil)
	if param.Granularity != "" {
		v := url.Values{}
		v.Add("granularity", param.Granularity)
		req.URL.RawQuery = v.Encode()
	}
	err = a.MakeRequest(req, &chart)

	return chart, err
}

func (a *Api) GetInstruments() (instrument *Instrument, err error) {
	req, _ := a.NewRequest("GET", fmt.Sprintf("v3/accounts/%s/instruments", a.AccountID), nil)
	err = a.MakeRequest(req, &instrument)

	return instrument, err
}
