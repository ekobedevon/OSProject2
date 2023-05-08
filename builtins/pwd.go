package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrsInvalidArgCount = errors.New("invalid argument count")
)

func PrintWorkingDirectory(args ...string) error {

	switch len(args) {
	case 0: // change to home directory if available
		CWD, err := os.Getwd()
		fmt.Println(CWD)
		return err
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)
	}
}
