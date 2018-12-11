package oanda

import (
	"fmt"
)

type OrderForm struct {
	Order struct {
		Type           string  `json:"type"`
		Instrument     string  `json:"instrument"`
		Units          float32 `json:"units"`
		Price          string  `json:"price"`
		TimeInForce    string  `json:"timeInForce"`
		PositionFill   string  `json:"positionFill"`
		StopLossOnFill struct {
			Price string `json:"price"`
		} `json:"stopLossOnFill"`
		TakeProfitOnFill struct {
			Price string `json:"price"`
		} `json:"takeProfitOnFill"`
	} `json:"order"`
}

type OrderCreatedResponse struct {
	OrderCreateTransaction struct {
		Type           string  `json:"type"`
		Instrument     string  `json:"instrument"`
		Units          string `json:"units"`
		Price          string  `json:"price"`
		TimeInForce      string `json:"timeInForce"`
		TriggerCondition string `json:"triggerCondition"`
		PartialFill      string `json:"partialFill"`
		PositionFill     string `json:"positionFill"`
		StopLossOnFill   struct {
			Price       string `json:"price"`
			TimeInForce string `json:"timeInForce"`
		} `json:"stopLossOnFill"`
		TrailingStopLossOnFill struct {
			Distance    string `json:"distance"`
			TimeInForce string `json:"timeInForce"`
		} `json:"trailingStopLossOnFill"`
		Reason    string `json:"reason"`
		ID        string `json:"id"`
		UserID    int    `json:"UserID"`
		AccountID string `json:"accountID"`
		BatchID   string `json:"batchID"`
		RequestID string `json:"requestID"`
		Time      string `json:"time"`
	} `json:"orderCreateTransaction"`
	OrderFillTransaction struct {
		Type           string  `json:"type"`
		Instrument     string  `json:"instrument"`
		Units          string `json:"units"`
		Price          string  `json:"price"`
		TradeOpened struct {
			Trade
		} `json:"tradeOpened"`
	} `json:"orderFillTransaction"`
	RelatedTransactionIDs []string `json:"relatedTransactionIDs"`
	LastTransactionID     string   `json:"lastTransactionID"`
}

type OrderCanceledResponse struct {
	OrderCancelTransaction struct {
		Type      string `json:"type"`
		OrderID   string `json:"orderID"`
		Reason    string `json:"reason"`
		ID        string `json:"id"`
		UserID    int    `json:"UserID"`
		AccountID string `json:"accountID"`
		BatchID   string `json:"batchID"`
		RequestID string `json:"requestID"`
		Time      string `json:"time"`
	} `json:"orderCancelTransaction"`
	RelatedTransactionIDs []string `json:"relatedTransactionIDs"`
	LastTransactionID     string   `json:"lastTransactionID"`
}

func (a *Api) PlaceOrder(form OrderForm) (res *OrderCreatedResponse, err error) {
	req, _ := a.NewRequest("POST", fmt.Sprintf("v3/accounts/%s/orders", a.AccountID), form)
	err = a.MakeRequest(req, &res)

	return res, err
}

func (a *Api) CancelOrder(orderID string) (res *OrderCanceledResponse, err error) {
	url := fmt.Sprintf("v3/accounts/%s/orders/%s/cancel", a.AccountID, orderID)
	req, _ := a.NewRequest("PUT", url, nil)
	err = a.MakeRequest(req, &res)

	return res, err
}
