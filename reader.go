// This package contains file reader logic
package main

import (
	"errors"
	"os"

	"gopkg.in/ini.v1"
)

// OSRelease contains parsed data from /etc/os-release files
type OSRelease struct {
	ID              string `json:"id"`
	IDLike          string `json:"id-like,omitempty"`
	Name            string `json:"name"`
	PrettyName      string `json:"pretty-name,omitempty"`
	Version         string `json:"version,omitempty"`
	VersionCodeName string `json:"version-codename,omitempty"`
}

func maybeReadINI(path string) (*OSRelease, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	section := cfg.Section("")
	release := OSRelease{
		ID:              section.Key("ID").String(),
		IDLike:          section.Key("ID_LIKE").String(),
		Name:            section.Key("NAME").String(),
		PrettyName:      section.Key("PRETTY_NAME").String(),
		VersionCodeName: section.Key("VERSION_CODENAME").String(),
		Version:         section.Key("VERSION_ID").String(),
	}
	return &release, nil
}
