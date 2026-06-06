package service

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
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

	dev := uinputUserDev{}
	copy(dev.Name[:], "Go Remote Mouse")
	dev.ID.Bustype, dev.ID.Vendor, dev.ID.Product, dev.ID.Version = 0x06, 0x1209, 0x5680, 1

	f.Write((*[unsafe.Sizeof(dev)]byte)(unsafe.Pointer(&dev))[:])
	uinputIoctl(f.Fd(), uiDevCreate, 0)

	mouseFile = f
	return nil
}

func MoveCursor(dx, dy float64) error {
	if err := initMouseDev(); err != nil {
		return err
	}

	x, y := int32(dx), int32(dy)

	if x != 0 {
		emitEvent(mouseFile, evRel, relX, x)
	}
	if y != 0 {
		emitEvent(mouseFile, evRel, relY, y)
	}
	if x != 0 || y != 0 {
		emitEvent(mouseFile, evSyn, 0, 0)
	}

	return nil
}

func Click(left bool) error {
	if err := initMouseDev(); err != nil {
		return err
	}

	btn := uint16(btnRight)
	if left {
		btn = btnLeft
	}

	emitEvent(mouseFile, evKey, btn, 1)
	emitEvent(mouseFile, evSyn, 0, 0)
	emitEvent(mouseFile, evKey, btn, 0)
	emitEvent(mouseFile, evSyn, 0, 0)

	return nil
}
