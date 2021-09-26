package utils

import (
	"net/http"
	"strings"
)

func GetId(r *http.Request) (id string) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}
