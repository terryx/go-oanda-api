package oanda

import (
	"fmt"
	"net/url"
)

type Pricing struct {
	Time string `json:"time"`
	Prices []struct {
		Type        string `json:"type"`
		Time        string `json:"time"`
		CloseoutBid string `json:"closeoutBid"`
		CloseoutAsk string `json:"closeoutAsk"`
		Status      string `json:"status"`
		Tradeable   bool   `json:"tradeable"`
		Instrument  string `json:"instrument"`
	} `json:"prices"`
}

func (a *Api) GetPricing(instrument string) (res *Pricing, err error) {
	req, _ := a.NewRequest("GET", fmt.Sprintf("v3/accounts/%s/pricing", a.AccountID), nil)
	v := url.Values{}
	v.Add("instruments", instrument)
	req.URL.RawQuery = v.Encode()

	err = a.MakeRequest(req, &res)

	return res, err
}