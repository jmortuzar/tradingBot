package alpaca

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/jmortuzar/tradingBot/internals/context"
	"github.com/shopspring/decimal"
)

func account() *alpaca.Account {
	acc, err := context.GetAlpacaClient().GetAccount()

	if err != nil {
		fmt.Println("Error getting account. App will exit.", err)
		os.Exit(-1)
	}

	return acc
}

func AccountEquity() decimal.Decimal {
	return account().Equity
}
