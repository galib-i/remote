package service

import (
	"fmt"
	"unsafe"
)

var (
	callSetCursorPos = newWinAPIHelper("SetCursorPos")
	callGetCursorPos = newWinAPIHelper("GetCursorPos")
	callMouseEvent   = newWinAPIHelper("mouse_event")
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

func mouseEvent(event uintptr) error {
	return callWinAPI(callMouseEvent, "mouse event failed", event, 0, 0, 0, 0)
}

func GetCursorPos() (float64, float64, error) {
	var pt POINT
	ret, _, err := callGetCursorPos(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, fmt.Errorf("failed to get cursor position: %w", err)
	}

	return float64(pt.X), float64(pt.Y), nil
}

func MoveCursor(x, y float64) error {
	return callWinAPI(callSetCursorPos, "failed to move cursor", uintptr(x), uintptr(y))
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

	if err := mouseEvent(downEvent); err != nil {
		return fmt.Errorf("failed to press mouse button: %w", err)
	}

	if err := mouseEvent(upEvent); err != nil {
		return fmt.Errorf("failed to release mouse button: %w", err)
	}

	return nil
}
