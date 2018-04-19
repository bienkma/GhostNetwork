package handler

import (
	"fmt"
	"strings"

	"net/http"
)

// IndexHandler handler home page
func IndexHandler(response http.ResponseWriter, request *http.Request) {
	ip := strings.Split(request.RemoteAddr, ":")[0]
	fmt.Fprintf(response, "Index visited by: %s", ip)
}
