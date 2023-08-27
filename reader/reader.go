// package reader contains ini file reader logic
package reader

import (
	"errors"
	"os"

	"github.com/oalders/is/types"
	"gopkg.in/ini.v1"
)

func MaybeReadINI(path string) (*types.OSRelease, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil //nolint:nilnil
		}
		return nil, errors.Join(errors.New("could not stat file"), err)
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return nil, errors.Join(errors.New("could not load file"), err)
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
