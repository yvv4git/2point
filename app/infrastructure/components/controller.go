package components

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sarulabs/di"
)

// Controller is a main controller entity.
type Controller struct {
	container di.Container
}

// SetContainer use for set parameter.
func (c *Controller) SetContainer(container di.Container) {
	c.container = container
}

// GetStorage passes pointer to inscance of Storage.
func (c *Controller) GetStorage() Storager {
	return c.container.Get("storage").(Storager)
}

// GetLogger passes pointer to instance of logger.
func (c *Controller) GetLogger() *log.Logger {
	return c.container.Get("logger").(*log.Logger)
}

// ServeHTTP is a callback for interface.
func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "From standart handler: %s", r.URL.Path)
}
