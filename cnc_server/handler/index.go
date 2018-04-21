package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// IndexHandler handler home page
func IndexHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Index visited by: %s", strings.Split(request.RemoteAddr, ":")[0])
}
