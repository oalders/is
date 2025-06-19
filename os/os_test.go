package os_test

import (
	"runtime"
	"testing"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	t.Parallel()
	tests := []string{"name", attr.Version}

	for _, v := range tests {
		ctx := types.Context{Debug: true}
		found, err := os.Info(&ctx, v)
		assert.NoError(t, err, v)
		assert.NotEmpty(t, found, v)
	}

	// id-like not present in Debian 11, so can't be in a blanket Linux test
	if runtime.GOOS == "linux" {
		tests := []string{"id", "pretty-name"}

		for _, v := range tests {
			ctx := types.Context{Debug: true}
			found, err := os.Info(&ctx, v)
			assert.NoError(t, err, v)
			assert.NotEmpty(t, found, v)
		}
	}
}
