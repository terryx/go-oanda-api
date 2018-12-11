## Getting Started
```
go get github.com/terryx/go-oanda-api
```

## Examples
```$xslt
import (
    "github.com/terryx/go-oanda-api/oanda"
)

# Initialize api
api := oanda.Api{
    ApiKey: "xxx",
    AccountID: "xxx-xxx",
    Endpoint: "https://api-fxpractice.oanda.com",
}

params := oanda.CandlestickParam{
    Name: "USD_CAD",
    Granularity: "H1",
}

# Get candlestick data
candle, _ := api.GetInstrumentCandles(params)
```