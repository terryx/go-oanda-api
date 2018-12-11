package oanda_test

import (
	"encoding/json"
	"fmt"
	"github.com/terryx/go-oanda-api/oanda"
	"log"
	"testing"
)

var api = oanda.Api{
	ApiKey: "your own api key",
	AccountID: "your account id",
	Endpoint: "https://api-fxpractice.oanda.com",
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func TestApi_GetInstrumentCandles(t *testing.T) {
	params := oanda.CandlestickParam{
		Name: "USD_CAD",
		Granularity: "H1",
	}
	candle, _ := api.GetInstrumentCandles(params)
	_ = PrettyPrint(candle)
}

func TestApi_PlaceOrder(t *testing.T) {
	req := oanda.OrderForm{}
	req.Order.Type = "LIMIT"
	req.Order.Instrument = "NAS100_USD"
	req.Order.Units = 1
	req.Order.Price = "5000.00"
	req.Order.TimeInForce = "GTC"
	req.Order.PositionFill = "DEFAULT"
	req.Order.StopLossOnFill.Price = "4000.00"
	req.Order.TakeProfitOnFill.Price = "9000.00"

	result, err := api.PlaceOrder(req)
	if err != nil {
		log.Print(err)
	}
	_ = PrettyPrint(result)

	cancelOrder, err := api.CancelOrder(result.OrderCreateTransaction.ID)
	if err != nil {
		log.Print(err)
	}
	_ = PrettyPrint(cancelOrder)
}

func TestApi_GetPricing(t *testing.T) {
	res, err := api.GetPricing("USD_CAD")
	if err != nil {
		log.Print(err)
	}
	_ = PrettyPrint(res)
}

func TestApi_CloseTrade(t *testing.T) {
	req := oanda.OrderForm{}
	req.Order.Type = "MARKET"
	req.Order.Instrument = "NZD_USD"
	req.Order.Units = 1
	req.Order.TimeInForce = "FOK"
	req.Order.PositionFill = "DEFAULT"
	req.Order.TakeProfitOnFill.Price = "1.90000"
	req.Order.StopLossOnFill.Price = "0.00009"

	newOrder, err := api.PlaceOrder(req)
	if err != nil {
		log.Print(err)
	}

	res, err := api.CloseTrade(newOrder.OrderFillTransaction.TradeOpened.TradeID)
	if err != nil {
		log.Print(err)
	}
	_ = PrettyPrint(res)
}

