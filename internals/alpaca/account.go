package alpaca

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/jmortuzar/tradingBot/internals/context"
	"github.com/shopspring/decimal"
)

// alpacaAccount returns the account details from Alpaca.
//
// If there is an error getting the account, the application will exit.
func alpacaAccount() *alpaca.Account {
	acc, err := context.GetAlpacaClient().GetAccount()

	if err != nil {
		fmt.Println("Error getting account. App will exit.")
		os.Exit(-1)
	}

	return acc
}

// AlpacaAccountEquity returns the account's equity as a decimal.Decimal.
func AlpacaAccountEquity() decimal.Decimal {
	return alpacaAccount().Equity
}
