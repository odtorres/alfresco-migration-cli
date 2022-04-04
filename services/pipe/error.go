package pipe

import (
	"fmt"
	"os"
)

//StopIfErrorReturnArg stop printing a message
func StopIfErrorReturnArg(i interface{}, err error) interface{} {

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return i
}

//StopIfErrorArg stop printing a message
func StopIfErrorArg(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//StopIfErrorReturn stop printing a message
func StopIfErrorReturn(f func() (interface{}, error)) interface{} {
	i, err := f()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return i
}

//StopIfError stop printing a message
func StopIfError(f func() error) {
	err := f()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
