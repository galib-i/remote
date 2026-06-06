package service

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	uiSetEvBit  = 0x40045564
	uiSetKeyBit = 0x40045565
	uiSetRelBit = 0x40045566
	uiDevCreate = 0x5501

	evSyn = 0x00
	evKey = 0x01
	evRel = 0x02
)

type uinputUserDev struct {
	Name                             [80]byte
	ID                               struct{ Bustype, Vendor, Product, Version uint16 }
	FfEffectsMax                     uint32
	Absmax, Absmin, Absfuzz, Absflat [64]int32
}

type inputEvent struct {
	Time       syscall.Timeval
	Type, Code uint16
	Value      int32
}

// uinputIoctl makes the syscall to configure the device
func uinputIoctl(fd uintptr, req, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, req, arg)
	if err != 0 {
		return err
	}
	return nil
}

// emitEvent writes the hardware action to the specific device file
func emitEvent(f *os.File, typ, code uint16, val int32) {
	ev := inputEvent{Type: typ, Code: code, Value: val}
	f.Write((*[unsafe.Sizeof(ev)]byte)(unsafe.Pointer(&ev))[:])
}
