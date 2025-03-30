// package version creates version objects from strings
package version

import (
	"fmt"

	goversion "github.com/hashicorp/go-version"
)

func NewVersion(vstring string) (*goversion.Version, error) {
	// func NewVersion(x string) (string,error) {`
	got, err := goversion.NewVersion(vstring)
	if err != nil {
		err = fmt.Errorf("parse version from \"%s\": %w", vstring, err)
	}
	return got, err
}
