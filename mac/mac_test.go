package mac_test

import (
	"runtime"
	"testing"

	"github.com/oalders/is/mac"
	"github.com/stretchr/testify/assert"
)

func TestCodeName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "ventura", mac.CodeName("13.1"))
	assert.Equal(t, "monterey", mac.CodeName("12.2"))
	assert.Equal(t, "big sur", mac.CodeName("11.2"))
	assert.Equal(t, "catalina", mac.CodeName("10.15"))
	assert.Equal(t, "mojave", mac.CodeName("10.14"))
	assert.Equal(t, "high sierra", mac.CodeName("10.13"))
	assert.Equal(t, "sierra", mac.CodeName("10.12"))
	assert.Equal(t, "el capitan", mac.CodeName("10.11"))
	assert.Equal(t, "yosemite", mac.CodeName("10.10"))
	assert.Equal(t, "mavericks", mac.CodeName("10.9"))
	assert.Equal(t, "mountain lion", mac.CodeName("10.8"))
	assert.Equal(t, "", mac.CodeName("10.7"))
	assert.Equal(t, "", mac.CodeName("9.0"))
	assert.Equal(t, "", mac.CodeName("-1"))
}

func TestVersion(t *testing.T) {
	t.Parallel()
	version, err := mac.Version()
	if runtime.GOOS == "darwin" {
		assert.NotEmpty(t, version)
		assert.NoError(t, err)
	} else {
		assert.Empty(t, version)
		assert.Error(t, err)
	}
}
