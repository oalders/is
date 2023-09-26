// package version creates verion objects from strings
package version

import (
	"errors"
	"fmt"

	goversion "github.com/hashicorp/go-version"
)

func NewVersion(vstring string) (*goversion.Version, error) {
	// func NewVersion(x string) (string,error) {`
	got, err := goversion.NewVersion(vstring)
	if err != nil {
		err = errors.Join(fmt.Errorf("parse version from \"%s\"", vstring), err)
	}
	return got, err
}
