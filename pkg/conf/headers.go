package conf

import (
	"net/http"
)

var (
	authHeaders = http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {""},
		"User-Agent":   {"RiotClient/79.0.2.1016.2013 rso-auth (Windows; 10;;Professional, x64)"},
	}
)

func SetAuthHeaders(headers http.Header) *http.Header {
	authHeaders = headers
	return &authHeaders
}

func GetAuthHeaders() *http.Header {
	return &authHeaders
}

func AddAuthHeaders(key, value string) *http.Header {
	authHeaders.Set(key, value)
	return &authHeaders
}
