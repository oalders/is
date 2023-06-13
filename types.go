// package main contains types used by this library
package main

type OSRelease struct {
	ID              string `json:"id"`
	IDLike          string `json:"id-like,omitempty"`
	Name            string `json:"name"`
	PrettyName      string `json:"pretty-name,omitempty"`
	Version         string `json:"version,omitempty"`
	VersionCodeName string `json:"version-codename,omitempty"`
}
