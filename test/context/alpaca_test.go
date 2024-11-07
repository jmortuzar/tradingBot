package context_test

import (
	"testing"

	"github.com/jmortuzar/tradingBot/internals/context"
)

func TestAlpacaClient(t *testing.T) {
	t.Run("Init And Get Alpaca Client", func(t *testing.T) {
		context.InitContext()
		alpacaClient := context.InitAlpacaClient()

		clientFromGlobalContext := context.GetAlpacaClient()

		if alpacaClient != clientFromGlobalContext {
			t.Errorf("Error configuring alpaca client. App will exit.")
		}

	})
}
