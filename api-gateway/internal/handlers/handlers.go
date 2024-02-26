package handlers

import (
	"fmt"
	"net/http"
)

// BaseHandler responds to the root URL
func BaseHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've reached the API gateway")
}
