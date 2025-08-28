//go:build freebsd || netbsd

// Package audio provides functions to read system audio volume information.
package audio

import (
	"errors"

	"github.com/oalders/is/types"
)

// VolumeInfo represents audio volume information.
type VolumeInfo struct {
	Level int  // Volume level (0-100)
	Muted bool // Whether audio is muted
}

// Summary retrieves the current system volume information.
// On FreeBSD and NetBSD, this returns an error as audio functionality is not supported.
func Summary(ctx *types.Context) (*VolumeInfo, error) {
	return nil, errors.New("audio functionality not supported on this platform")
}

// Level returns just the volume level (0-100).
// On FreeBSD and NetBSD, this returns an error as audio functionality is not supported.
func Level() (int, error) {
	return 0, errors.New("audio functionality not supported on this platform")
}

// IsMuted returns whether the system audio is currently muted.
// On FreeBSD and NetBSD, this returns an error as audio functionality is not supported.
func IsMuted() (bool, error) {
	return false, errors.New("audio functionality not supported on this platform")
}
