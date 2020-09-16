package router

import (
	"2point/app/infrastructure/services"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMux(t *testing.T) {
	container := services.NewApp()
	mux := NewMux(container)

	assert.IsType(t, &http.ServeMux{}, mux)
}
