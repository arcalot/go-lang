package catch

import (
	"go.arcalot.io/lang/try"
)

// Error creates a handler that catches any error types. This is typically used as a last error handler.
func Error(f func(err error)) try.CatchHandler {
	return func(e any) bool {
		err, ok := e.(error)
		if !ok {
			return false
		}
		f(err)
		return true
	}
}
