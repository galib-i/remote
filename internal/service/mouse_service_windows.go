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

func getCursorPos() (float64, float64, error) {
	var pt POINT
	ret, _, err := callGetCursorPos(uintptr(unsafe.Pointer(&pt)))
	if ret == 0 {
		return 0, 0, fmt.Errorf("failed to get cursor position: %w", err)
	}

	return float64(pt.X), float64(pt.Y), nil
}

func MoveCursor(dx, dy float64) error {
	currentX, currentY, err := getCursorPos()
	if err != nil {
		return err
	}

	return callWinAPI(callSetCursorPos, "failed to move cursor", uintptr(currentX+dx), uintptr(currentY+dy))
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

	var hardwareErr error
	emit := func(event uintptr) {
		if hardwareErr == nil {
			hardwareErr = mouseEvent(event)
		}
	}

	emit(downEvent)
	emit(upEvent)

	return hardwareErr
}
