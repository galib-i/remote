package service

import (
	"fmt"
	"strings"
)

var (
	callKeybd_event = newWinAPIHelper("keybd_event")
	procVkKeyScan   = user32.NewProc("VkKeyScanA")
)

const (
	KEYEVENTF_KEYUP = 0x0002
	SHIFT_MOD       = 0x01
	VK_SHIFT        = 0x10
)

var specialKeys = map[string]byte{
	"backspace": 0x08, "enter": 0x0D,
}

func keybdEvent(virtualKey byte, flags uintptr) error {
	return callWinAPI(callKeybd_event, "keyboard event failed", uintptr(virtualKey), 0, flags, 0)
}

func getVirtualKeyAndModifiers(key rune) (byte, byte, error) {
	ret, _, _ := procVkKeyScan.Call(uintptr(key))
	vk := byte(ret & 0xFF)
	mod := byte((ret >> 8) & 0xFF)

	if vk == 0xFF {
		return 0, 0, fmt.Errorf("unsupported key: %c", key)
	}

	return vk, mod, nil
}

func PressKey(key string) error {
	if len(key) == 0 {
		return fmt.Errorf("cannot press empty key")
	}

	// Check for special keys first
	if specialKey, exists := specialKeys[strings.ToLower(key)]; exists {
		if err := keybdEvent(specialKey, 0); err != nil {
			return fmt.Errorf("failed to press %s down: %w", key, err)
		}

		if err := keybdEvent(specialKey, KEYEVENTF_KEYUP); err != nil {
			return fmt.Errorf("failed to release %s: %w", key, err)
		}

		return nil
	}

	r := []rune(key)[0]
	virtualKey, mod, err := getVirtualKeyAndModifiers(r)
	if err != nil {
		return err
	}

	var hardwareErr error
	emit := func(vk byte, flags uintptr) {
		if hardwareErr == nil {
			hardwareErr = keybdEvent(vk, flags)
		}
	}

	// Press shift if needed
	if mod&SHIFT_MOD != 0 {
		emit(VK_SHIFT, 0)
	}

	emit(virtualKey, 0)               // Key down
	emit(virtualKey, KEYEVENTF_KEYUP) // Key up

	// Release modifiers
	if mod&SHIFT_MOD != 0 {
		emit(VK_SHIFT, KEYEVENTF_KEYUP)
	}

	return hardwareErr
}
