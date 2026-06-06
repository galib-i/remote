package service

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

const keyLeftShift = 42

var keyMap = map[string]uint16{
	"enter": 28, "backspace": 14, " ": 57,
	"a": 30, "b": 48, "c": 46, "d": 32, "e": 18, "f": 33, "g": 34, "h": 35,
	"i": 23, "j": 36, "k": 37, "l": 38, "m": 50, "n": 49, "o": 24, "p": 25,
	"q": 16, "r": 19, "s": 31, "t": 20, "u": 22, "v": 47, "w": 17, "x": 45,
	"y": 21, "z": 44, "1": 2, "2": 3, "3": 4, "4": 5, "5": 6, "6": 7,
	"7": 8, "8": 9, "9": 10, "0": 11, ",": 51, ".": 52, "/": 53,
}

var kbdFile *os.File

func initKbdDev() error {
	if kbdFile != nil {
		return nil
	}

	f, err := os.OpenFile("/dev/uinput", os.O_WRONLY|syscall.O_NONBLOCK, 0220)
	if err != nil {
		return fmt.Errorf("failed to open /dev/uinput: %w", err)
	}

	uinputIoctl(f.Fd(), uiSetEvBit, evKey)
	for i := uint16(1); i <= 120; i++ {
		uinputIoctl(f.Fd(), uiSetKeyBit, uintptr(i))
	}

	dev := uinputUserDev{}
	copy(dev.Name[:], "Go Remote Keyboard")
	dev.ID.Bustype, dev.ID.Vendor, dev.ID.Product, dev.ID.Version = 0x06, 0x1209, 0x5679, 1

	f.Write((*[unsafe.Sizeof(dev)]byte)(unsafe.Pointer(&dev))[:])
	uinputIoctl(f.Fd(), uiDevCreate, 0)

	kbdFile = f
	return nil
}

func PressKey(keyStr string) error {
	if err := initKbdDev(); err != nil || len(keyStr) == 0 {
		return err
	}

	isCapital := len(keyStr) == 1 && keyStr[0] >= 'A' && keyStr[0] <= 'Z'
	keyCode, exists := keyMap[strings.ToLower(keyStr)]
	if !exists {
		return fmt.Errorf("unsupported key: %s", keyStr)
	}

	if isCapital {
		emitEvent(kbdFile, evKey, keyLeftShift, 1)
		emitEvent(kbdFile, evKey, keyCode, 1)
		emitEvent(kbdFile, evSyn, 0, 0)

		emitEvent(kbdFile, evKey, keyCode, 0)
		emitEvent(kbdFile, evKey, keyLeftShift, 0)
		emitEvent(kbdFile, evSyn, 0, 0)
	} else {
		emitEvent(kbdFile, evKey, keyCode, 1)
		emitEvent(kbdFile, evSyn, 0, 0)
		emitEvent(kbdFile, evKey, keyCode, 0)
		emitEvent(kbdFile, evSyn, 0, 0)
	}

	return nil
}
