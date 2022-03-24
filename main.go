package main

import (
	"net/http"

	"github.com/tislais/go-webserver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
