package sigctx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// WithSignal returns a context that is canceled when the specified signal is received.
func WithSignal(sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancelCause(context.Background())
	received := make(chan os.Signal, 1)
	signal.Notify(received, sig...)

	go func() {
		sig := <-received
		cause := fmt.Errorf("received signal: %s: %w", sig, context.Canceled)
		cancel(cause)
	}()

	return ctx
}
