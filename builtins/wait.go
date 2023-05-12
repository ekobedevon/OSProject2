// Implements shell builtin "wait"
// tat0154
package builtins

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"
)

var (
	errMissArgs  = errors.New("missing arguments")
	errInvalArgs = errors.New("invalid arguments")
	procExists   bool
)

func BuiltinWait(args ...string) error {

	//Check for args
	if len(args) < 1 {
		return fmt.Errorf("%w: unexpected arguments", errMissArgs)
	}

	for i := range args {
		argPid, convErr := strconv.Atoi(args[i])
		if convErr != nil {
			return fmt.Errorf("%w: failed to convert pid", errInvalArgs)
		}

		if argPid <= 0 {
			fmt.Printf("Invalid PID %d. Skipping.\n", argPid)
			continue
		}

		procExists = true //Assume process exists
		fmt.Printf("Waiting on Process %d... ", argPid)

		for procExists { //Waiting Loop

			//OS Specific Logic required due to Signal differences
			if runtime.GOOS == "windows" { //Windows Systems
				waitProc := exec.Command("TASKLIST", "/FI", fmt.Sprintf("PID eq %d", argPid))
				output, err := waitProc.Output()
				if err != nil { //Unable to run TASKLIST command
					procExists = false
					fmt.Printf("Unable to fetch PID list. ")
				}

				//Checks the output for the phrase "INFO: no tasks", stores inverted bool result
				//Contains the phrase (true) -> Process does not exist (false)
				//Doesn't contain the phrase (false) -> Process found (true)
				procExists = !bytes.Contains(output, []byte("INFO: No tasks"))

			} else { //UNIX Systems
				waitProc, err := os.FindProcess(argPid)
				if err != nil { //If FindProcess fails, process doesn't exists
					procExists = false
					fmt.Printf("Unable to find PID %d. ", argPid)
				} else {
					err := waitProc.Signal(syscall.Signal(0)) //Send SIG 0
					if err != nil {                           //if Signal returned an error, assume the process no longer exists
						procExists = false
					}
					errno, _ := err.(syscall.Errno) //Get Error Number
					if errno == syscall.EPERM {     //Check if error was access denied, if so -> process still exists
						procExists = true
					}
				}

			}

		}

		fmt.Printf("Done waiting on PID %d! \n", argPid)

	}

	return nil

}
