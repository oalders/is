package reader_test

import (
	"context"
	"testing"

	"github.com/oalders/is/reader"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

func TestMaybeReadINI(t *testing.T) {
	t.Parallel()
	ctx := types.Context{
		Context: context.Background(),
	}
	{
		release, err := reader.MaybeReadINI(&ctx, "../testdata/etc/os-release")
		assert.NoError(t, err)
		assert.Equal(t, "18.04", release.Version)
	}
	{
		// if the file does not exist on this system, that's not an error
		release, err := reader.MaybeReadINI(&ctx, "../testdata/etc/os-releasezzz")
		assert.NoError(t, err)
		assert.Nil(t, release)
	}
	{
		ctx.Debug = true
		// if the file cannot be parsed, that's an error
		release, err := reader.MaybeReadINI(&ctx, "../testdata/etc/not-an-ini-file")
		assert.Error(t, err)
		assert.Nil(t, release)
	}
}
