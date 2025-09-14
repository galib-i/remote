package service

import (
	"github.com/itchyny/volume-go"
)

func VolumeUp() error {
	vol, err := volume.GetVolume()
	if err != nil {
		return err
	}

	err = volume.SetVolume(vol + 2)
	if err != nil {
		return err
	}

	return nil
}

func VolumeDown() error {
	vol, err := volume.GetVolume()
	if err != nil {
		return err
	}

	err = volume.SetVolume(vol - 2)
	if err != nil {
		return err
	}

	return nil
}

func MuteVolume() error {
	status, err := volume.GetMuted()
	if err != nil {
		return err
	}

	if status != true {
		volume.Mute()
	}

	return nil
}

func UnmuteVolume() error {
	status, err := volume.GetMuted()
	if err != nil {
		return err
	}

	if status == true {
		volume.Unmute()
	}

	return nil
}
