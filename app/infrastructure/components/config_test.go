package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigSimple(t *testing.T) {
	storage, err := NewStorage("simple")

	assert.Nil(t, err)
	assert.IsType(t, &StorageSimple{}, storage)
}

func TestConfigBad(t *testing.T) {
	storage, err := NewStorage("bads")

	assert.Nil(t, storage)
	assert.Equal(t, "Unknown type", err.Error())
}
