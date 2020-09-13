package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"2point/app/infrastructure/services"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCount(t *testing.T) {
	var countController CountController

	container := services.NewApp()
	countController.SetContainer(container)

	// Create web server for tests.
	server := httptest.NewServer(&countController)
	defer server.Close()

	// Do request.
	resp, err := http.Get(server.URL)
	t.Log(server.URL)

	// Check.
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
