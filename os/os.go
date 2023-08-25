// Package os handles OS info parsing
package os

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/oalders/is/mac"
	"github.com/oalders/is/reader"
	"github.com/oalders/is/types"
)

const (
	darwin        = "darwin"
	name          = "name"
	linux         = "linux"
	osReleaseFile = "/etc/os-release"
)

func Info(ctx *types.Context, argName string) (string, error) {
	result := ""
	switch argName {
	case "id":
		if runtime.GOOS == linux {
			if ctx.Debug {
				log.Println("Trying to parse " + osReleaseFile)
			}
			release, err := reader.MaybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.ID != "" {
				result = release.ID
			}
		}
	case "id-like":
		if runtime.GOOS == linux {
			if ctx.Debug {
				log.Println("Trying to parse " + osReleaseFile)
			}
			release, err := reader.MaybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.IDLike != "" {
				result = release.IDLike
			}
		}
	case "pretty-name":
		if runtime.GOOS == linux {
			if ctx.Debug {
				log.Println("Trying to parse " + osReleaseFile)
			}
			release, err := reader.MaybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.PrettyName != "" {
				result = release.PrettyName
			}
		}
	case name:
		result = runtime.GOOS
	case "version":
		if runtime.GOOS == darwin {
			o, err := mac.Version()
			if err != nil {
				return result, err
			}
			result = o
		} else if runtime.GOOS == linux {
			if ctx.Debug {
				log.Println("Trying to parse " + osReleaseFile)
			}
			release, err := reader.MaybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.Version != "" {
				result = release.Version
			}
		}
	case "version-codename":
		if runtime.GOOS == linux {
			if ctx.Debug {
				log.Println("Trying to parse " + osReleaseFile)
			}
			release, err := reader.MaybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.VersionCodeName != "" {
				result = release.VersionCodeName
			}
		} else if runtime.GOOS == darwin {
			o, err := mac.Version()
			if err != nil {
				return result, err
			}
			name := mac.CodeName(o)
			if name != "" {
				result = name
			}
		}
	}
	if result != "" {
		ctx.Success = true
	}

	return result, nil
}

func Aggregated() (string, error) {
	release, err := reader.MaybeReadINI(osReleaseFile)
	if err != nil {
		return "", err
	}
	if release == nil {
		release = &types.OSRelease{}
	}
	release.Name = runtime.GOOS

	if runtime.GOOS == darwin {
		v, err := mac.Version()
		if err != nil {
			return "", err
		}
		release.Version = v
		release.VersionCodeName = mac.CodeName(release.Version)
	}
	data, err := json.MarshalIndent(release, "", "    ")
	if err != nil {
		return "", errors.Join(fmt.Errorf("could not marshal indented JSON (%+v)", release), err)
	}

	return string(data), nil
}
