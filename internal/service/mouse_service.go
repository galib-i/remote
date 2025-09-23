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
	procMouseEvent   = user32.NewProc("mouse_event")
)

type POINT struct {
	X, Y int32
}

const (
	MOUSEEVENTF_LEFTDOWN  = 0x0002
	MOUSEEVENTF_LEFTUP    = 0x0004
	MOUSEEVENTF_RIGHTDOWN = 0x0008
	MOUSEEVENTF_RIGHTUP   = 0x0010
)

func GetCursorPos() (float64, float64, error) {
	var pt POINT
	ret, _, err := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, fmt.Errorf("failed to get cursor position: %w", err)
	}

	return float64(pt.X), float64(pt.Y), nil
}

func MoveCursor(x, y float64) error {
	ret, _, err := procSetCursorPos.Call(
		uintptr(x),
		uintptr(y),
	)
	if ret == 0 {
		return fmt.Errorf("failed to move cursor: %w", err)
	}

	return nil
}

func Click(left bool) error {
	var downEvent, upEvent uintptr

	if left {
		downEvent = MOUSEEVENTF_LEFTDOWN
		upEvent = MOUSEEVENTF_LEFTUP
	} else {
		downEvent = MOUSEEVENTF_RIGHTDOWN
		upEvent = MOUSEEVENTF_RIGHTUP
	}

	// Mouse button down
	ret, _, err := procMouseEvent.Call(downEvent, 0, 0, 0, 0)
	if ret == 0 {
		return fmt.Errorf("failed to press mouse button: %w", err)
	}

	// Mouse button up
	ret, _, err = procMouseEvent.Call(upEvent, 0, 0, 0, 0)
	if ret == 0 {
		return fmt.Errorf("failed to release mouse button: %w", err)
	}

	return nil
}
