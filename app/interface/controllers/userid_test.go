package controllers

import (
	"2point/app/infrastructure/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleUserId(t *testing.T) {
	var useridController UseridController

	container := services.NewApp()
	useridController.SetContainer(container)

	// Create request.
	req, err := http.NewRequest("GET", "/?user_id=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create recorder.
	rr := httptest.NewRecorder()
	handler := http.Handler(&useridController)

	handler.ServeHTTP(rr, req)

	// Check
	assert.Equal(t, http.StatusOK, rr.Code)
}
