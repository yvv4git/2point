package services

import (
	"2point/app/infrastructure/components"
	"log"
	"testing"

	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
)

func TestContainer(t *testing.T) {
	container := NewApp()

	assert.Implements(t, (*di.Container)(nil), container)

	logger := container.Get("logger").(*log.Logger)
	config := container.Get("config").(*components.Config)

	assert.IsType(t, &log.Logger{}, logger)
	assert.IsType(t, &components.Config{}, config)
}
