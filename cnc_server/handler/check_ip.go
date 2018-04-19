package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// IPHandler function provice check ip public method
func IPHandler(response http.ResponseWriter, request *http.Request) {
	ip := strings.Split(request.RemoteAddr, ":")[0]
	fmt.Fprintf(response, "ip: %s", ip)
}
