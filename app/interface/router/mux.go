package router

import (
	"net/http"

	"2point/app/interface/controllers"

	"github.com/sarulabs/di"
)

// GetMux is a simple method for router logic encapsulation.
func NewMux(container di.Container) *http.ServeMux {
	var countController controllers.CountController
	var useridController controllers.UseridController

	countController.SetContainer(container)
	useridController.SetContainer(container)

	mux := http.NewServeMux()
	mux.Handle("/count", &countController)
	mux.Handle("/", &useridController)

	return mux
}
