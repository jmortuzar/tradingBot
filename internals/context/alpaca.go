package context

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
)

const AlpacaContextName string = "alpacaClient"

func InitAlpacaClient() *alpaca.Client {
	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    os.Getenv("ALPACA_API_KEY"),
		APISecret: os.Getenv("ALPACA_API_SECRET"),
		BaseURL:   os.Getenv("ALPACA_BASE_URL"),
	})

	GlobalContext.PutValue(AlpacaContextName, client)
	return client
}

func GetAlpacaClient() *alpaca.Client {
	client, ok := GlobalContext.GetValue(AlpacaContextName)
	if !ok {
		fmt.Println("Error getting alpaca client. App will exit.")
		os.Exit(-1)
	}
	return client.(*alpaca.Client)
}
