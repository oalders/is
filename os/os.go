// Package os handles OS info parsing
package os

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/mac"
	"github.com/oalders/is/reader"
	"github.com/oalders/is/types"
)

const (
	darwin        = "darwin"
	linux         = "linux"
	osReleaseFile = "/etc/os-release"
)

func Info(ctx *types.Context, argName string) (string, error) {
	maybeDebug(ctx)
	if argName == attr.Name {
		return runtime.GOOS, nil
	}

	if runtime.GOOS == linux {
		return linuxOS(ctx, argName)
	}

	if runtime.GOOS != darwin {
		return "", nil
	}

	macVersion, err := mac.Version()
	if err != nil {
		return "", err
	}

	switch argName {
	case attr.Version:
		return macVersion, nil
	case attr.VersionCodename:
		return mac.CodeName(macVersion), nil
	}

	return "", nil
}

func AsJSON(ctx *types.Context) (string, error) {
	release, err := reader.MaybeReadINI(ctx, osReleaseFile)
	if err != nil {
		return "", err
	}
	if release == nil {
		release = &types.OSRelease{}
	}
	release.Name = runtime.GOOS

	if runtime.GOOS == darwin {
		v, versionErr := mac.Version()
		if versionErr != nil {
			return "", versionErr
		}
		release.Version = v
		release.VersionCodeName = mac.CodeName(release.Version)
	}
	data, err := json.MarshalIndent(release, "", "    ")
	if err != nil {
		return "", fmt.Errorf("could not marshal indented JSON (%+v): %w", release, err)
	}

	return string(data), nil
}

func linuxOS(ctx *types.Context, argName string) (string, error) {
	release, err := reader.MaybeReadINI(ctx, osReleaseFile)
	if err != nil {
		return "", err
	}
	if release == nil {
		return "", errors.New("release info cannot be found")
	}

	result := ""
	switch argName {
	case "id":
		result = release.ID
	case "id-like":
		result = release.IDLike
	case "pretty-name":
		result = release.PrettyName
	case attr.Version:
		result = release.Version
	case attr.VersionCodename:
		result = release.VersionCodeName
	}
	if result != "" {
		return result, nil
	}
	return "", nil
}

func maybeDebug(ctx *types.Context) {
	if ctx.Debug {
		log.Printf(
			"Run %q to see available os data\n",
			"is known summary os",
		)
	}
}
