package service

import (
	"fmt"
	"os"
	"syscall"
)

const (
	relX     = 0x00
	relY     = 0x01
	btnLeft  = 0x110
	btnRight = 0x111
)

var mouseFile *os.File

func initMouseDev() error {
	if mouseFile != nil {
		return nil
	}

	f, err := os.OpenFile("/dev/uinput", os.O_WRONLY|syscall.O_NONBLOCK, 0220)
	if err != nil {
		return fmt.Errorf("failed to open /dev/uinput: %w", err)
	}

	uinputIoctl(f.Fd(), uiSetEvBit, evKey)
	uinputIoctl(f.Fd(), uiSetEvBit, evRel)

	uinputIoctl(f.Fd(), uiSetKeyBit, btnLeft)
	uinputIoctl(f.Fd(), uiSetKeyBit, btnRight)

	uinputIoctl(f.Fd(), uiSetRelBit, relX)
	uinputIoctl(f.Fd(), uiSetRelBit, relY)

	if err := registerUinputDevice(f, "Go-Remote Mouse"); err != nil {
		return err
	}

	mouseFile = f
	return nil
}

func MoveCursor(dx, dy float64) error {
	if err := initMouseDev(); err != nil {
		return err
	}

	x, y := int32(dx), int32(dy)

	var hardwareErr error
	emit := func(typ, code uint16, val int32) {
		if hardwareErr == nil {
			hardwareErr = emitEvent(mouseFile, typ, code, val)
		}
	}

	if x != 0 {
		emit(evRel, relX, x)
	}
	if y != 0 {
		emit(evRel, relY, y)
	}
	if x != 0 || y != 0 {
		emit(evSyn, 0, 0)
	}

	return hardwareErr
}

func Click(left bool) error {
	if err := initMouseDev(); err != nil {
		return err
	}

	btn := uint16(btnRight)
	if left {
		btn = btnLeft
	}

	var hardwareErr error
	emit := func(typ, code uint16, val int32) {
		if hardwareErr == nil {
			hardwareErr = emitEvent(mouseFile, typ, code, val)
		}
	}

	emit(evKey, btn, 1)
	emit(evSyn, 0, 0)
	emit(evKey, btn, 0)
	emit(evSyn, 0, 0)

	return hardwareErr
}
