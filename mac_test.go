package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMacCodeName(t *testing.T) {
	assert.Equal(t, "ventura", macCodeName("13.1"))
	assert.Equal(t, "monterey", macCodeName("12.2"))
	assert.Equal(t, "big sur", macCodeName("11.2"))
	assert.Equal(t, "catalina", macCodeName("10.15"))
	assert.Equal(t, "mojave", macCodeName("10.14"))
	assert.Equal(t, "high sierra", macCodeName("10.13"))
	assert.Equal(t, "sierra", macCodeName("10.12"))
	assert.Equal(t, "el capitan", macCodeName("10.11"))
	assert.Equal(t, "yosemite", macCodeName("10.10"))
	assert.Equal(t, "mavericks", macCodeName("10.9"))
	assert.Equal(t, "mountain lion", macCodeName("10.8"))
	assert.Equal(t, "", macCodeName("10.7"))
}

func TestMacVersion(t *testing.T) {
	version, err := macVersion()
	if runtime.GOOS == "darwin" {
		assert.NotEmpty(t, version)
		assert.NoError(t, err)
	} else {
		assert.Empty(t, version)
		assert.Error(t, err)
	}
}
