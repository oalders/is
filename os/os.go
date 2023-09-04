// Package os handles OS info parsing
package os

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"

	"github.com/oalders/is/attr"
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
	if argName == "name" {
		ctx.Success = true
		return runtime.GOOS, nil
	}

	if runtime.GOOS == linux {
		return linuxOS(ctx, argName)
	}

	if runtime.GOOS != "darwin" {
		return "", nil
	}

	result := ""
	macVersion, err := mac.Version()
	if err != nil {
		return result, err
	}
	switch argName {
	case "version":
		result = macVersion
	case "version-codename":
		name := mac.CodeName(macVersion)
		if name != "" {
			result = name
		}
	}
	if result != "" {
		ctx.Success = true
	}

	return result, nil
}

func Aggregated(ctx *types.Context) (string, error) {
	release, err := reader.MaybeReadINI(ctx, osReleaseFile)
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
	case "version-codename":
		result = release.VersionCodeName
	}
	if result != "" {
		ctx.Success = true
	}
	return result, err
}
