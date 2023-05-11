package builtins

import (
	"errors"
	"fmt"
	"io/ioutil"
)

var(
	ErrArgCount =errors.New("argument count invalid")
)

func ListDirectory(args ...string) error {
	if len(args) > 0 {  // if arguemnt is provided, displays message
		return fmt.Errorf("%w: expected zero arguments", ErrArgCount)
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {  // if there is still errors, displays message
		return fmt.Errorf("error, cant read directory %w", err)
	}

	for _, file := range files {  // prints the full list if no errors
		fmt.Println(file.Name())
	}

	return nil
}
