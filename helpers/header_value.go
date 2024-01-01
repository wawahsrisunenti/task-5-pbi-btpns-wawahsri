package helpers

import (
	"net/http"
)

func GetAuthorizationHeader(r *http.Request) string {
	return r.Header.Get("Authorization")
}
