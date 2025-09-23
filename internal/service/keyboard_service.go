package service

import (
	"fmt"
	"strings"
)

var (
	callKeybd_event      = newWinAPIHelper("keybd_event")
	callGetAsyncKeyState = newWinAPIHelper("GetAsyncKeyState")
)

const (
	KEYEVENTF_KEYUP = 0x0002
)

func keybdEvent(virtualKey byte, flags uintptr) error {
	return callWinAPI(callKeybd_event, "keyboard event failed", uintptr(virtualKey), 0, flags, 0)
}

// Virtual key codes for letters A-Z
var letterKeys = map[string]byte{
	"a": 0x41, "b": 0x42, "c": 0x43, "d": 0x44, "e": 0x45,
	"f": 0x46, "g": 0x47, "h": 0x48, "i": 0x49, "j": 0x4A,
	"k": 0x4B, "l": 0x4C, "m": 0x4D, "n": 0x4E, "o": 0x4F,
	"p": 0x50, "q": 0x51, "r": 0x52, "s": 0x53, "t": 0x54,
	"u": 0x55, "v": 0x56, "w": 0x57, "x": 0x58, "y": 0x59,
	"z": 0x5A,
}

func PressKey(key string) error {
	key = strings.ToLower(key)

	virtualKey, exists := letterKeys[key]
	if !exists {
		return fmt.Errorf("unsupported key: %s", key)
	}

	// Key down
	if err := keybdEvent(virtualKey, 0); err != nil {
		return fmt.Errorf("failed to press key down: %w", err)
	}

	// Key up
	if err := keybdEvent(virtualKey, KEYEVENTF_KEYUP); err != nil {
		return fmt.Errorf("failed to release key: %w", err)
	}

	return nil
}

func IsKeyPressed(key string) (bool, error) {
	key = strings.ToLower(key)

	virtualKey, exists := letterKeys[key]
	if !exists {
		return false, fmt.Errorf("unsupported key: %s", key)
	}

	ret, _, err := callGetAsyncKeyState(
		uintptr(virtualKey),
	)
	// Note: GetAsyncKeyState can return 0 when not failing, so we don't use callWinAPI here.
	if err != nil && err.Error() != "The operation completed successfully." {
		return false, fmt.Errorf("failed to get key state: %w", err)
	}

	// The high-order bit is 1 if the key is down
	return (ret & 0x8000) != 0, nil
}
