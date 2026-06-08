package service

import (
	"fmt"

	"github.com/itchyny/volume-go"
)

const volumeStep = 2

func AdjustVolume(increase bool) error {
	vol, err := volume.GetVolume()
	if err != nil {
		return fmt.Errorf("failed to get current volume: %w", err)
	}

	adjustment := volumeStep
	if !increase {
		adjustment = -volumeStep
	}

	newVol := vol + adjustment
	if newVol < 0 {
		newVol = 0
	} else if newVol > 100 {
		newVol = 100
	}

	if err := volume.SetVolume(newVol); err != nil {
		return fmt.Errorf("failed to set new volume to %d: %w", newVol, err)
	}

	return nil
}

func ToggleMute() error {
	mute, err := volume.GetMuted()
	if err != nil {
		return fmt.Errorf("failed to get current mute state: %w", err)
	}

	if mute {
		if err := volume.Unmute(); err != nil {
			return fmt.Errorf("failed to unmute volume: %w", err)
		}
		return nil
	}

	if err := volume.Mute(); err != nil {
		return fmt.Errorf("failed to mute volume: %w", err)
	}

	return nil
}
