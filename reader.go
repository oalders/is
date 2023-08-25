// This package contains file reader logic
package main

import (
	"errors"
	"os"

	"github.com/oalders/is/types"
	"gopkg.in/ini.v1"
)

func maybeReadINI(path string) (*types.OSRelease, error) {
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
	release := types.OSRelease{
		ID:              section.Key("ID").String(),
		IDLike:          section.Key("ID_LIKE").String(),
		Name:            section.Key("NAME").String(),
		PrettyName:      section.Key("PRETTY_NAME").String(),
		VersionCodeName: section.Key("VERSION_CODENAME").String(),
		Version:         section.Key("VERSION_ID").String(),
	}
	return &release, nil
}
