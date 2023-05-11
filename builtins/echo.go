// Implements shell builtin "echo"
// tat0154
package builtins

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errInvalidArgs  = errors.New("invalid arguments")
	errMissingArgs  = errors.New("missing arguments")
	enableBackslash bool
	trailingNewline bool
	startIndex      int
)

func Echo(args ...string) error {

	//Check for args
	if len(args) < 1 {
		return fmt.Errorf("%w: unexpected arguments", errMissingArgs)
	}

	//OUTPUT OPTIONS
	//Set Defaults
	enableBackslash = false
	trailingNewline = true

	//Check if specified options are set
	hasOptions := strings.HasPrefix(args[0], "-")

	if hasOptions { //Option Declared
		option := strings.TrimPrefix(args[0], "-")

		switch option {
		case "e":
			enableBackslash = true
		case "E": //Default options
			break
		case "ne":
			trailingNewline = false
			enableBackslash = true
		case "nE":
			trailingNewline = false
		case "n":
			trailingNewline = false
		default:
			return fmt.Errorf("%w: unexpected option", errInvalidArgs)
		}
	}

	//PRINT ARGUMENTS

	//Setup and check args
	if hasOptions {
		startIndex = 1
		if len(args) < 2 {
			return fmt.Errorf("%w: unexpected arguments", errMissingArgs)
		}
	} else {
		startIndex = 0
		if len(args) < 1 {
			return fmt.Errorf("%w: unexpected arguments", errMissingArgs)
		}
	}

	//Print loop
	for i := startIndex; i < len(args); i++ {
		if enableBackslash {
			modStr, err := strconv.Unquote(`"` + args[i] + `"`)
			if err != nil {
				return fmt.Errorf("%w: invalid arguments", errInvalidArgs)
			}
			fmt.Print(modStr)
		} else {

			fmt.Print(args[i])
		}

		fmt.Print(" ") //Add space between arguments
	}

	if trailingNewline {
		fmt.Printf("\n")
	}

	return nil

}
