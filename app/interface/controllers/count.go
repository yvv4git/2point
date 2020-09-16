package controllers

import (
	"fmt"
	"net/http"

	"2point/app/infrastructure/components"
)

// CountController is a handler for count request.
type CountController struct {
	components.Controller
}

// ServeHTTP is a callback for interface.
func (c *CountController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	storage := c.GetStorage()
	count := storage.GetValue()
	logger := c.GetLogger()

	logger.Println("request to count")
	fmt.Fprintf(w, "Count value: %d\n", count)
}
