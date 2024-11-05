package context_test

import (
	"testing"

	"github.com/jmortuzar/tradingBot/internals/context"
)

func TestContext(t *testing.T) {
	t.Run("Create Context", func(t *testing.T) {
		ctx := context.InitContext()
		ctxAux := context.GlobalContext

		if ctx != ctxAux {
			t.Errorf("Error creating context")
		}

	})

	t.Run("Put and Get values", func(t *testing.T) {
		ctx := createCtx(t)

		expect := "value"

		ctxVal := *ctx

		ctxVal.PutValue("key", expect)

		value, ok := ctxVal.GetValue("key")

		if !ok || value != expect {
			t.Errorf("No value foun or wrong value. %v", value)
		}
	})
}

func createCtx(t testing.TB) *context.Context {
	t.Helper()
	ctx := context.InitContext()
	ctxAux := context.GlobalContext

	if ctx != ctxAux {
		t.Errorf("Error creating context")
	}

	return ctx
}
