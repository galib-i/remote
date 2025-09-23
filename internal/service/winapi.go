package service

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var (
	user32 = windows.NewLazySystemDLL("user32.dll")
)

// newWinAPIHelper creates a reusable helper function for a specific WinAPI procedure.
func newWinAPIHelper(procName string) func(...uintptr) (uintptr, uintptr, error) {
	proc := user32.NewProc(procName)
	return func(args ...uintptr) (uintptr, uintptr, error) {
		return proc.Call(args...)
	}
}

// callWinAPI is a generic function to execute a call to a Windows API procedure.
// It checks for a return value of 0 as a failure condition.
func callWinAPI(apiFunc func(...uintptr) (uintptr, uintptr, error), failureMsg string, args ...uintptr) error {
	ret, _, err := apiFunc(args...)
	if ret == 0 {
		return fmt.Errorf("%s: %w", failureMsg, err)
	}
	return nil
}
