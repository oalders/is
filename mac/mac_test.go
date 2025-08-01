package mac_test

import (
	"context"
	"runtime"
	"testing"

	"github.com/oalders/is/mac"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestCodeName(t *testing.T) {
	t.Parallel()
	tests := [][]string{
		{"13.1", "ventura"},
		{"12.2", "monterey"},
		{"11.2", "big sur"},
		{"10.15", "catalina"},
		{"10.14", "mojave"},
		{"10.13", "high sierra"},
		{"10.12", "sierra"},
		{"10.11", "el capitan"},
		{"10.10", "yosemite"},
		{"10.9", "mavericks"},
		{"10.8", "mountain lion"},
		{"10.7", ""},
		{"9.0", ""},
		{"-1", ""},
	}

	for _, v := range tests {
		assert.Equal(t, v[1], mac.CodeName(v[0]))
	}
}

func TestVersion(t *testing.T) {
	t.Parallel()
	ctx := types.Context{Context: context.Background()}
	version, err := mac.Version(&ctx)
	if runtime.GOOS == "darwin" {
		assert.NotEmpty(t, version)
		assert.NoError(t, err)
	} else {
		assert.Empty(t, version)
		assert.Error(t, err)
	}
}
