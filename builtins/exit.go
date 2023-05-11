package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrsInvalidArgCount = errors.New("invalid arguments")
)

func ExitCommand(args ...string) error {

	switch len(args) {
	case 0: 
	  os.Exit(0)
	default:
		return fmt.Errorf("%w: unexpected arguments", ErrInvalidArgCount)
	}
}
