// Package audio provides functions to read system audio volume information.
package audio

import (
	"fmt"

	"github.com/itchyny/volume-go"
	"github.com/oalders/is/types"
)

// VolumeInfo represents audio volume information.
type VolumeInfo struct {
	Level int  // Volume level (0-100)
	Muted bool // Whether audio is muted
}

// Summary retrieves the current system volume information.
func Summary(ctx *types.Context) (*VolumeInfo, error) {
	vol, err := Level()
	if err != nil {
		return nil, err
	}

	muted, err := IsMuted()
	if err != nil {
		return nil, err
	}
	ctx.Success = true

	return &VolumeInfo{
		Level: vol,
		Muted: muted,
	}, nil
}

// Level returns just the volume level (0-100).
func Level() (int, error) {
	level, err := volume.GetVolume()
	if err != nil {
		return 0, fmt.Errorf("get level: %w", err)
	}
	return level, nil
}

// IsMuted returns whether the system audio is currently muted.
func IsMuted() (bool, error) {
	isMuted, err := volume.GetMuted()
	if err != nil {
		return false, fmt.Errorf("get mute status: %w", err)
	}
	return isMuted, nil
}
