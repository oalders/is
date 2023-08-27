package version_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/oalders/is/version"
)

func TestCompareCLIVersions(t *testing.T) {
	t.Parallel()
	_, err := version.NewVersion("3.3")
	assert.NoError(t, err)

	_, err = version.NewVersion("x")
	assert.Error(t, err)
}
