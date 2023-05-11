// Implements shell builtin "echo"
// tat0154
package builtins

import (
	"errors"
	"fmt"
)

var (
	invalidArgs = errors.New("invalid arguments")
)

func echo(args ...string) error {

	switch len(args) {
	case 0: //No Argmunets
		return fmt.Errorf("%w: unexpected arguments", invalidArgs)

	default:
		return fmt.Errorf("%w: unexpected arguments", invalidArgs)
	}

}
