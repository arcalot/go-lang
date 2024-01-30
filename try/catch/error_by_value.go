package catch

import (
	"errors"

	"go.arcalot.io/lang/try"
)

// ErrorByValue creates an error handler that catches fixed value errors. For example, errors are typically
// declared outside any functions like this:
//
//	var ErrMyError = fmt.Errorf("some value")
func ErrorByValue(value error, f func(err error)) try.CatchHandler {
	return func(e any) bool {
		err, ok := e.(error)
		if !ok {
			return false
		}
		if errors.Is(err, value) {
			f(err)
			return true
		}
		return false
	}
}
