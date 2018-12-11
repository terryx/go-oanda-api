package oanda

import (
	"fmt"
)

type Account struct {
	Account struct {
		ID      string `json:"id"`
		Balance string `json:"balance"`
		Orders  []struct {
			ID               string `json:"id"`
			CreateTime       string `json:"createTime"`
			Type             string `json:"type"`
			TradeID          string `json:"tradeID"`
			Price            string `json:"price"`
			TimeInForce      string `json:"timeInForce"`
			TriggerCondition string `json:"triggerCondition"`
			State            string `json:"state"`
		} `json:"orders"`
		Trades []struct {
			ID                    string `json:"id"`
			Instrument            string `json:"instrument"`
			Price                 string `json:"price"`
			OpenTime              string `json:"openTime"`
			InitialUnits          string `json:"initialUnits"`
			InitialMarginRequired string `json:"initialMarginRequired"`
			State                 string `json:"state"`
			CurrentUnits          string `json:"currentUnits"`
			RealizedPL            string `json:"realizedPL"`
			Financing             string `json:"financing"`
			TakeProfitOrderID     string `json:"takeProfitOrderID"`
			StopLossOrderID       string `json:"stopLossOrderID"`
			UnrealizedPL          string `json:"unrealizedPL"`
			MarginUsed            string `json:"marginUsed"`
		}
	} `json:"account"`
}

func (a *Api) GetAccount() (account *Account, err error) {
	req, _ := a.NewRequest("GET", fmt.Sprintf("v3/accounts/%s", a.AccountID), nil)
	err = a.MakeRequest(req, &account)

	return account, err
}
