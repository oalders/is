// package reader contains ini file reader logic
package reader

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/oalders/is/types"
	"gopkg.in/ini.v1"
)

func MaybeReadINI(ctx *types.Context, path string) (*types.OSRelease, error) {
	if ctx.Debug {
		log.Println("Trying to parse " + path)
	}
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil //nolint:nilnil
		}
		return nil, fmt.Errorf("could not stat file: %w", err)
	}

	cfg, err := ini.Load(path)
	if err != nil {
		return nil, fmt.Errorf("could not load file: %w", err)
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
