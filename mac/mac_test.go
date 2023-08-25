package mac

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "ventura", CodeName("13.1"))
	assert.Equal(t, "monterey", CodeName("12.2"))
	assert.Equal(t, "big sur", CodeName("11.2"))
	assert.Equal(t, "catalina", CodeName("10.15"))
	assert.Equal(t, "mojave", CodeName("10.14"))
	assert.Equal(t, "high sierra", CodeName("10.13"))
	assert.Equal(t, "sierra", CodeName("10.12"))
	assert.Equal(t, "el capitan", CodeName("10.11"))
	assert.Equal(t, "yosemite", CodeName("10.10"))
	assert.Equal(t, "mavericks", CodeName("10.9"))
	assert.Equal(t, "mountain lion", CodeName("10.8"))
	assert.Equal(t, "", CodeName("10.7"))
}

func TestVersion(t *testing.T) {
	t.Parallel()
	version, err := Version()
	if runtime.GOOS == "darwin" {
		assert.NotEmpty(t, version)
		assert.NoError(t, err)
	} else {
		assert.Empty(t, version)
		assert.Error(t, err)
	}
}
