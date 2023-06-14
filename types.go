// package main contains types used by this library
package main

type osRelease struct {
	Arch            string `json:"arch"`
	ID              string `json:"id,omitempty"`
	IDLike          string `json:"id-like,omitempty"`
	Name            string `json:"name"`
	PrettyName      string `json:"pretty-name,omitempty"`
	Version         string `json:"version,omitempty"`
	VersionCodeName string `json:"version-codename,omitempty"`
}
