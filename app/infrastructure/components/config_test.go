package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/2endpoint/app/infrastructure/components"
)

func TestConfigSimple(t *testing.T) {
	storage, err := components.NewStorage("simple")

	assert.Nil(t, err)
	assert.IsType(t, &StorageSimple{}, storage)
}

func TestConfigBad(t *testing.T) {
	storage, err := components.NewStorage("bads")

	assert.Nil(t, storage)
	assert.Equal(t, "Unknown type", err.Error())
}
