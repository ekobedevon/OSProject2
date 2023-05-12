// Implements shell builtin "wait"
// tat0154
package builtins

import (
	"errors"
	"fmt"
)

var (
	errMissArgs = errors.New("missing arguments")
)

func BuiltinWait(args ...string) error {

	//Check for args
	if len(args) < 1 {
		return fmt.Errorf("%w: unexpected arguments", errMissArgs)
	}

	return nil

}
