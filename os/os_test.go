package os_test

import (
	"context"
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

	for _, attr := range tests {
		ctx := &types.Context{
			Context: context.Background(),
			Debug:   true,
		}
		found, err := os.Info(ctx, attr)
		assert.NoError(t, err, attr)
		assert.NotEmpty(t, found, attr)
	}

	// id-like not present in Debian 11, so can't be in a blanket Linux test
	if runtime.GOOS == "linux" {
		tests := []string{"id", "pretty-name"}

		for _, attr := range tests {
			ctx := &types.Context{
				Context: context.Background(),
				Debug:   true,
			}
			found, err := os.Info(ctx, attr)
			assert.NoError(t, err, attr)
			assert.NotEmpty(t, found, attr)
		}
	}
}
