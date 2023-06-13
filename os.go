// This file handles OS info parsing
package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

const osReleaseFile = "/etc/os-release"

func osInfo(ctx *Context, argName string) (string, error) {
	result := ""
	switch argName {
	case "arch":
		result = runtime.GOARCH
	case "id":
		if runtime.GOOS == "linux" {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.ID != "" {
				result = release.ID
			}
		}
	case "id-like":
		if runtime.GOOS == "linux" {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.IDLike != "" {
				result = release.IDLike
			}
		}
	case "pretty-name":
		if runtime.GOOS == "linux" {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.PrettyName != "" {
				result = release.PrettyName
			}
		}
	case "name":
		result = runtime.GOOS
	case "version":
		if runtime.GOOS == "darwin" {
			o, err := macVersion()
			if err != nil {
				return result, err
			}
			result = o
		} else if runtime.GOOS == "linux" {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.Version != "" {
				result = release.Version
			}
		}
	case "version-codename":
		if runtime.GOOS == "linux" {
			if ctx.Debug {
				fmt.Println("Trying to parse " + osReleaseFile)
			}
			release, err := maybeReadINI(osReleaseFile)
			if err == nil && release != nil && release.VersionCodeName != "" {
				result = release.VersionCodeName
			}
		} else if runtime.GOOS == "darwin" {
			o, err := macVersion()
			if err != nil {
				return result, err
			}
			name := macCodeName(o)
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

func aggregatedOS() (string, error) {
	release, err := maybeReadINI(osReleaseFile)
	if err != nil {
		return "", err
	}
	release.Name = runtime.GOOS

	if runtime.GOOS == "darwin" {
		v, err := macVersion()
		if err != nil {
			return "", err
		}
		release.Version = v
	}
	data, err := json.MarshalIndent(release, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
