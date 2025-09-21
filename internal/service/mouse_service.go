package service

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	procSetCursorPos = user32.NewProc("SetCursorPos")
	procGetCursorPos = user32.NewProc("GetCursorPos")
)

type POINT struct {
	X, Y int32
}

func GetCursorPos() (float64, float64, error) {
	var pt POINT
	ret, _, err := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, fmt.Errorf("failed to get cursor position: %w", err)
	}
	return float64(pt.X), float64(pt.Y), nil
}

func MoveMouse(x, y float64) error {
	ret, _, err := procSetCursorPos.Call(
		uintptr(x),
		uintptr(y),
	)
	if ret == 0 {
		return fmt.Errorf("failed to move cursor: %w", err)
	}
	return nil
}
