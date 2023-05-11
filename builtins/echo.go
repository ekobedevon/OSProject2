// Implements shell builtin "echo"
// tat0154
package builtins

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errInvalidArgs = errors.New("invalid arguments")
	errMissingArgs = errors.New("missing arguments")
)

func Echo(args ...string) error {

	//Check for args
	if len(args) < 1 {
		return fmt.Errorf("%w: unexpected arguments", errMissingArgs)
	}

	//Check for declared options
	hasOptions := strings.HasPrefix(args[0], "-")

	if hasOptions { //Option Declared
		option := strings.TrimPrefix(args[0], "-")

		switch option {
		case "n":
			fmt.Print("-n specified")
		case "e":
			fmt.Print("-e specified")
		case "E":
			fmt.Print("-E specified")
		case "ne":
			fmt.Print("-ne specified")
		case "nE":
			fmt.Print("-nE specified")
		default:
			return fmt.Errorf("%w: unexpected option", errInvalidArgs)
		}

	} else {
		return fmt.Errorf("%w: unexpected arguments", errInvalidArgs)
	}

	return fmt.Errorf("end of func")

}
