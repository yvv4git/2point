package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"2point/app/infrastructure/components"
)

// UseridController is a handler for userid request.
type UseridController struct {
	components.Controller
}

// ServeHTTP is a callback for interface.
func (c *UseridController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := c.GetLogger()
	storage := c.GetStorage()

	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if nil != err {
		logger.Println(err)
	}

	storage.Increment(userId)

	logger.Printf("request to user with userId: %d \n", userId)
	fmt.Fprintf(w, "User userId: %d \n", userId)
}
