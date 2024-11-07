package context

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
)

func InitAlpacaClient() *alpaca.Client {
	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    os.Getenv("ALPACA_API_KEY"),
		APISecret: os.Getenv("ALPACA_API_SECRET"),
		BaseURL:   os.Getenv("ALPACA_BASE_URL"),
	})

	GlobalContext.PutValue("alpacaClient", client)
	return client
}

func GetAlpacaClient() *alpaca.Client {
	client, ok := GlobalContext.GetValue("alpacaClient")
	if !ok {
		fmt.Println("Error getting alpaca client. App will exit.")
		os.Exit(-1)
	}
	return client.(*alpaca.Client)
}
