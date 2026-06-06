package service

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const (
	uiSetEvBit  = 0x40045564
	uiSetKeyBit = 0x40045565
	uiDevCreate = 0x5501

	evSyn = 0x00
	evKey = 0x01
)

const (
	keyLeftShift = 42
)

var keyMap = map[string]uint16{
	"enter": 28, "backspace": 14, " ": 57,
	"a": 30, "b": 48, "c": 46, "d": 32, "e": 18, "f": 33, "g": 34, "h": 35,
	"i": 23, "j": 36, "k": 37, "l": 38, "m": 50, "n": 49, "o": 24, "p": 25,
	"q": 16, "r": 19, "s": 31, "t": 20, "u": 22, "v": 47, "w": 17, "x": 45,
	"y": 21, "z": 44, "1": 2, "2": 3, "3": 4, "4": 5, "5": 6, "6": 7,
	"7": 8, "8": 9, "9": 10, "0": 11, ",": 51, ".": 52, "/": 53,
}

type kbdUinputUserDev struct {
	Name                             [80]byte
	ID                               struct{ Bustype, Vendor, Product, Version uint16 }
	FfEffectsMax                     uint32
	Absmax, Absmin, Absfuzz, Absflat [64]int32
}

type kbdInputEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}

var kbdFile *os.File

func initKbdDev() error {
	if kbdFile != nil {
		return nil
	}

	f, err := os.OpenFile("/dev/uinput", os.O_WRONLY|syscall.O_NONBLOCK, 0220)
	if err != nil {
		return fmt.Errorf("failed to open /dev/uinput (must run with sudo): %w", err)
	}

	// Enable event types
	kbdIoctl(f.Fd(), uiSetEvBit, evKey)

	// Enable all basic keyboard keys
	for i := uint16(1); i <= 120; i++ {
		kbdIoctl(f.Fd(), uiSetKeyBit, uintptr(i))
	}

	// Simulate a virtual keyboard
	dev := kbdUinputUserDev{}
	copy(dev.Name[:], "Go Remote Keyboard")
	dev.ID.Bustype = 0x06
	dev.ID.Vendor = 0x1234
	dev.ID.Product = 0x1234
	dev.ID.Version = 1

	b := (*[unsafe.Sizeof(dev)]byte)(unsafe.Pointer(&dev))
	if _, err := f.Write(b[:]); err != nil {
		f.Close()
		return fmt.Errorf("failed to write dev info: %w", err)
	}

	if err := kbdIoctl(f.Fd(), uiDevCreate, 0); err != nil {
		f.Close()
		return fmt.Errorf("failed to create uinput keyboard device: %w", err)
	}
	kbdFile = f
	return nil
}

func kbdIoctl(fd, req, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, req, arg)
	if err != 0 {
		return err
	}
	return nil
}

func kbdEmit(fd *os.File, typ, code uint16, val int32) error {
	ev := kbdInputEvent{
		Type:  typ,
		Code:  code,
		Value: val,
	}
	b := (*[unsafe.Sizeof(ev)]byte)(unsafe.Pointer(&ev))
	_, err := fd.Write(b[:])
	return err
}

func PressKey(keyStr string) error {
	if err := initKbdDev(); err != nil || len(keyStr) == 0 {
		return err
	}

	// Detect if the incoming character is a capital letter
	isCapital := len(keyStr) == 1 && keyStr[0] >= 'A' && keyStr[0] <= 'Z'

	// Match with key map
	lowerStr := strings.ToLower(keyStr)
	keyCode, exists := keyMap[lowerStr]
	if !exists {
		return fmt.Errorf("unsupported key: %s", keyStr)
	}

	// If capital, hold Shift down
	if isCapital {
		kbdEmit(kbdFile, evKey, keyLeftShift, 1) // Shift down
		kbdEmit(kbdFile, evKey, keyCode, 1)      // Letter down
		kbdEmit(kbdFile, evSyn, 0, 0)

		// Release both at the same time
		kbdEmit(kbdFile, evKey, keyCode, 0)      // Letter pp
		kbdEmit(kbdFile, evKey, keyLeftShift, 0) // Shift pp
		kbdEmit(kbdFile, evSyn, 0, 0)
	} else {
		// Lowercase letter
		kbdEmit(kbdFile, evKey, keyCode, 1)
		kbdEmit(kbdFile, evSyn, 0, 0)
		kbdEmit(kbdFile, evKey, keyCode, 0)
		kbdEmit(kbdFile, evSyn, 0, 0)
	}

	return nil
}
