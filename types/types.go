package types

// Context type tracks top level debugging flag.
type Context struct {
	Debug   bool
	Success bool
}

//nolint:tagliatelle
type OSRelease struct {
	ID              string `json:"id,omitempty"`
	IDLike          string `json:"id-like,omitempty"`
	Name            string `json:"name"`
	PrettyName      string `json:"pretty-name,omitempty"`
	Version         string `json:"version,omitempty"`
	VersionCodeName string `json:"version-codename,omitempty"`
}
