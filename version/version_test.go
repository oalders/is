package version_test

import (
	"testing"

	"github.com/oalders/is/version"
	"github.com/stretchr/testify/assert"
)

func TestCompareCLIVersions(t *testing.T) {
	t.Parallel()
	_, err := version.NewVersion("3.3")
	assert.NoError(t, err)

	_, err = version.NewVersion("x")
	assert.Error(t, err)
}
