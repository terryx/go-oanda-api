package oanda

import "fmt"

type Trade struct {
	Price   string  `json:"price"`
	TradeID string  `json:"tradeID"`
	Units   string `json:"units"`
}

type TradeResponse struct {
	OrderFillTransaction struct {
		TradesClosed []struct {
			Trade
			RealizedPL             string `json:"realizedPL"`
			Financing              string `json:"financing"`
			GuaranteedExecutionFee string `json:"guaranteedExecutionFee"`
			HalfSpreadCost         string `json:"halfSpreadCost"`
		} `json:"tradesClosed"`
	} `json:"orderFillTransaction"`
}

func (a *Api) CloseTrade(tradeID string) (res *TradeResponse, err error) {
	url := fmt.Sprintf("v3/accounts/%s/trades/%s/close", a.AccountID, tradeID)
	req, _ := a.NewRequest("PUT", url, nil)
	err = a.MakeRequest(req, &res)

	return res, err
}
