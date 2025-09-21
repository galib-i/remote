package service

import (
	"github.com/itchyny/volume-go"
)

const volumeStep = 2

func AdjustVolume(increase bool) error {
	vol, err := volume.GetVolume()
	if err != nil {
		return err
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

	return volume.SetVolume(newVol)
}

func ToggleMute() error {
	mute, err := volume.GetMuted()
	if err != nil {
		return err
	}

	if mute {
		return volume.Unmute()
	}

	return volume.Mute()
}
